package gogi

import (
	"github.com/teris-io/shortid"
)

type MessageModel struct {
	User User
	Room Room
	Data interface{}
}

type Room interface {
	Init()
	GetName() string
	SetName(string)
	GetShortID() string
	SetShortID()
	AddUser(User)
	RemoveUser(string)
	GetUsers() map[string]User
	SetUsers(map[string]User)
	GetHost() string
	SetHost(string)
}

type RoomStatus int

const (
	Waiting RoomStatus = iota + 1
	Playing
	Finished
)

type RoomModel struct {
	ShortID string
	Name    string
	Status  RoomStatus
	Users   map[string]User
	Host    string
}

func (rm *RoomModel) Init() {
	rm.Users = make(map[string]User)
	rm.SetShortID()
	rm.Name = "room-" + rm.GetShortID()
}

func (rm *RoomModel) GetName() string {
	return rm.Name
}

func (rm *RoomModel) SetName(s string) {
	rm.Name = s
}

func (rm *RoomModel) GetShortID() string {
	return rm.ShortID
}

func (rm *RoomModel) SetShortID() {
	rm.ShortID, _ = shortid.Generate()
}

func (rm *RoomModel) AddUser(u User) {
	rm.Users[u.GetShortID()] = u
}

func (rm *RoomModel) RemoveUser(sid string) {
	delete(rm.Users, sid)
}

func (rm *RoomModel) GetUsers() map[string]User {
	return rm.Users
}

func (rm *RoomModel) SetUsers(u map[string]User) {
	rm.Users = u
}

func (rm *RoomModel) SetHost(s string) {
	rm.Host = s
}

func (rm *RoomModel) GetHost() string {
	return rm.Host
}

type Manager interface {
	Events()
	HandleEvent(string, MessageModel)
	Register(string, func(MessageModel))
}

type ManagerModel struct {
	EventHandlers map[string]func(MessageModel)
}

func (mm *ManagerModel) Register(name string, f func(MessageModel)) {
	mm.EventHandlers[name] = f
}

func (mm *ManagerModel) HandleEvent(event string, m MessageModel) {
	mm.EventHandlers[event](m)
}
