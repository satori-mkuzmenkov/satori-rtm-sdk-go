package connection

import (
	"encoding/json"
	"github.com/satori-com/satori-rtm-sdk-go/logger"
	"github.com/satori-com/satori-rtm-sdk-go/rtm/pdu"
	"golang.org/x/net/websocket"
	"math"
	"strconv"
	"sync"
	"time"
)

const (
	MAX_ID                     = math.MaxInt32
	MAX_ACKS_QUEUE_LENGTH      = 10000
	MAX_UNPROCESSED_ACKS_QUEUE = 100
)

type Connection struct {
	wsConn *websocket.Conn
	lastID int
	acks   acksType
	mutex  sync.Mutex
}

type acksType struct {
	ch        chan pdu.RTMQuery
	listeners map[string]chan pdu.RTMQuery
	mutex     sync.Mutex
}

func New(endpoint string) (*Connection, error) {
	var err error

	conn := &Connection{}
	conn.lastID = 0
	conn.wsConn, err = websocket.Dial(endpoint, "", "http://localhost")
	if err != nil {
		return nil, err
	}

	conn.initAcks()

	return conn, nil
}

func (c *Connection) Close() {
	defer func() {
		// Channel can be already closed. Call recover to avoid panic when closing closed channel
		recover()
	}()
	if c.wsConn != nil {
		c.wsConn.Close()
	}

	// Close Ack listeners channel
	for _, ch := range c.acks.listeners {
		close(ch)
	}

	close(c.acks.ch)
}

func (c *Connection) SendAck(action string, body json.RawMessage) (<-chan pdu.RTMQuery, error) {
	query := pdu.RTMQuery{
		Action: action,
		Body:   body,
		Id:     c.nextID(),
	}

	ch := make(chan pdu.RTMQuery, 1)
	c.addListener(query.Id, ch)

	return ch, c.socketSend(query)
}

func (c *Connection) Send(action string, body json.RawMessage) error {
	query := pdu.RTMQuery{
		Action: action,
		Body:   body,
	}

	return c.socketSend(query)
}

func (c *Connection) socketSend(query pdu.RTMQuery) error {
	message, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	logger.Debug("send>", string(message))
	_, err = c.wsConn.Write(message)

	if err != nil {
		c.Close()
		return err
	}

	return nil
}

func (c *Connection) Read() (pdu.RTMQuery, error) {
	var response pdu.RTMQuery

	d := json.NewDecoder(c.wsConn)
	err := d.Decode(&response)
	if err != nil {
		c.Close()
		return pdu.RTMQuery{}, err
	}

	logger.Debug("recv<", response.String())

	if len(response.Id) != 0 {
		c.acks.ch <- response
	}

	return response, nil
}

func (c *Connection) SetDeadline(t time.Time) {
	c.wsConn.SetDeadline(t)
}

func (c *Connection) nextID() string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.lastID == MAX_ID {
		c.lastID = 0
	}
	c.lastID++
	return strconv.Itoa(c.lastID)
}

func (c *Connection) initAcks() {
	c.acks.ch = make(chan pdu.RTMQuery, MAX_UNPROCESSED_ACKS_QUEUE)
	c.acks.listeners = make(map[string]chan pdu.RTMQuery, MAX_ACKS_QUEUE_LENGTH)

	go func(c *Connection) {
		for response := range c.acks.ch {
			c.acks.mutex.Lock()
			ch := c.acks.listeners[response.Id]
			c.acks.mutex.Unlock()

			// Exception for the "search" API: Do not delete listener channel until the last message
			if response.Action != "rtm/search/data" {
				c.deleteListener(response.Id)
			}

			if ch != nil {
				if pdu.GetResponseCode(response) != pdu.CODE_DATA_REQUEST {
					defer close(ch)
				}
				ch <- response
			}
		}
	}(c)
}

func (c *Connection) addListener(id string, channel chan pdu.RTMQuery) {
	c.acks.mutex.Lock()
	defer c.acks.mutex.Unlock()
	c.acks.listeners[id] = channel
}

func (c *Connection) deleteListener(id string) {
	c.acks.mutex.Lock()
	defer c.acks.mutex.Unlock()
	delete(c.acks.listeners, id)
}