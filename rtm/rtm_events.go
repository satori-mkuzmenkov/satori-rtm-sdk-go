package rtm

const (
	EVENT_STOPPED          = "enterStopped"
	EVENT_LEAVE_STOPPED    = "leaveStopped"
	EVENT_CONNECTING       = "enterConnecting"
	EVENT_LEAVE_CONNECTING = "leaveConnecting"
	EVENT_CONNECTED        = "enterConnected"
	EVENT_LEAVE_CONNECTED  = "leaveConnected"
	EVENT_AWAITING         = "enterAwaiting"
	EVENT_LEAVE_AWAITING   = "leaveAwaiting"
	EVENT_START            = "start"
	EVENT_STOP             = "stop"
	EVENT_OPEN             = "open"
	EVENT_CLOSE            = "close"
	EVENT_ERROR            = "error"
	EVENT_DATA_ERROR       = "dataError"
	EVENT_AUTHENTICATED    = "authenticated"
)

// EVENT_STOPPED

func (rtm *RTM) OnStopped(callback func()) interface{} {
	return rtm.On(EVENT_STOPPED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnStoppedOnce(callback func()) {
	rtm.Once(EVENT_STOPPED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveStopped(callback func()) interface{} {
	return rtm.On(EVENT_LEAVE_STOPPED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveStoppedOnce(callback func()) {
	rtm.Once(EVENT_LEAVE_STOPPED, func(data interface{}) {
		callback()
	})
}

// EVENT_CONNECTING

func (rtm *RTM) OnConnecting(callback func()) interface{} {
	return rtm.On(EVENT_CONNECTING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnConnectingOnce(callback func()) {
	rtm.Once(EVENT_CONNECTING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveConnecting(callback func()) interface{} {
	return rtm.On(EVENT_LEAVE_CONNECTING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveConnectingOnce(callback func()) {
	rtm.Once(EVENT_LEAVE_CONNECTING, func(data interface{}) {
		callback()
	})
}

// EVENT_CONNECTED

func (rtm *RTM) OnConnected(callback func()) interface{} {
	return rtm.On(EVENT_CONNECTED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnConnectedOnce(callback func()) {
	rtm.Once(EVENT_CONNECTED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveConnected(callback func()) interface{} {
	return rtm.On(EVENT_LEAVE_CONNECTED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveConnectedOnce(callback func()) {
	rtm.Once(EVENT_LEAVE_CONNECTED, func(data interface{}) {
		callback()
	})
}

// EVENT_AWAITING

func (rtm *RTM) OnAwaiting(callback func()) interface{} {
	return rtm.On(EVENT_AWAITING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnAwaitingOnce(callback func()) {
	rtm.Once(EVENT_AWAITING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveAwaiting(callback func()) interface{} {
	return rtm.On(EVENT_LEAVE_AWAITING, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnLeaveAwaitingOnce(callback func()) {
	rtm.Once(EVENT_LEAVE_AWAITING, func(data interface{}) {
		callback()
	})
}

// Other events

func (rtm *RTM) OnStart(callback func()) interface{} {
	return rtm.On(EVENT_START, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnStartOnce(callback func()) {
	rtm.Once(EVENT_START, func(data interface{}) {
		callback()
	})
}

func (rtm *RTM) OnStop(callback func()) interface{} {
	return rtm.On(EVENT_STOP, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnStopOnce(callback func()) {
	rtm.Once(EVENT_STOP, func(data interface{}) {
		callback()
	})
}

func (rtm *RTM) OnOpen(callback func()) interface{} {
	return rtm.On(EVENT_OPEN, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnOpenOnce(callback func()) {
	rtm.Once(EVENT_OPEN, func(data interface{}) {
		callback()
	})
}

func (rtm *RTM) OnError(callback func(err RTMError)) interface{} {
	return rtm.On(EVENT_ERROR, func(data interface{}) {
		err := data.(RTMError)
		callback(err)
	})
}
func (rtm *RTM) OnErrorOnce(callback func(err RTMError)) {
	rtm.Once(EVENT_ERROR, func(data interface{}) {
		err := data.(RTMError)
		callback(err)
	})
}

func (rtm *RTM) OnDataError(callback func(err RTMError)) interface{} {
	return rtm.On(EVENT_DATA_ERROR, func(data interface{}) {
		err := data.(RTMError)
		callback(err)
	})
}
func (rtm *RTM) OnDataErrorOnce(callback func(err RTMError)) {
	rtm.Once(EVENT_DATA_ERROR, func(data interface{}) {
		err := data.(RTMError)
		callback(err)
	})
}

func (rtm *RTM) OnAuthenticated(callback func()) interface{} {
	return rtm.On(EVENT_AUTHENTICATED, func(data interface{}) {
		callback()
	})
}
func (rtm *RTM) OnAuthenticatedOnce(callback func()) {
	rtm.Once(EVENT_AUTHENTICATED, func(data interface{}) {
		callback()
	})
}
