package gogi

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (g *Game) roomNewHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		u := 10
		user := g.Context.NewUser()
		err := g.Context.DB.First(user, u).Error
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		room := g.Context.NewRoom()
		room.Init()
		room.AddUser(user)
		g.Context.Rooms[room.GetShortID()] = room
		http.Redirect(w, r, g.Context.Prefix+"/room/id/"+room.GetShortID(), 302)
	}
}
