package gogi

type State interface {
	Events()
	Register(string, func(Message))
}

type StateModel struct {
	EventHandlers map[string]func(Message)
}

func (sm *StateModel) Register(name string, f func(Message)) {
	sm.EventHandlers[name] = f
}
