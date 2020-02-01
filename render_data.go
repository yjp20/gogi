package gogi

type RenderData struct {
	Context *Context
	Self    User
	User    User
	Users   []User
	Room    Room
	Rooms   []Room
}
