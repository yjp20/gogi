package gogi

import (
	"html/template"
	"log"
	"net/http"

	// "github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

const guestAuthTemplate = `<div class="card-content">
	<form class="form" method="post" action="{{.Context.Prefix}}/auth/guest">
		<div class="field">
			<label class="label"> Guest Login </label>
			<div class="control">
				<input name="nick" class="input" type="text" placeholder="Nickname">
			</div>
		</div>
		<div class="field">
			<button class="button is-fullwidth"> Go </button>
		</div>
	</form>
</div>`

type GuestAuth struct {
	TemplateParsed *template.Template
}

func (a *GuestAuth) Init(c *Context) error {
	var err error
	if c.DB == nil {
		log.Fatal("Guest auth requires a DB provider")
	}
	if a.TemplateParsed == nil {
		a.TemplateParsed = template.New("")
		_, err = a.TemplateParsed.Parse(guestAuthTemplate)
	}
	return err
}

func (a *GuestAuth) Template() *template.Template {
	return a.TemplateParsed
}

func (a *GuestAuth) Handlers(c *Context) []AuthMethodHandler {
	nickHandler := AuthMethodHandler{
		"POST", "/auth/guest", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			nick := r.Form.Get("nick")
			u := c.NewUser()
			u.SetNick(nick)
			u.SetShortID()
			u.SetTemp(true)
			err = c.DB.Save(u).Error
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			s, _ := c.Store.Get(r, "session")
			s.Options.MaxAge = 24 * 60 * 60 * 1000 // Save for one day
			s.Values["self"] = u.GetID()
			err = s.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			http.Redirect(w, r, c.Prefix+"/rooms", 302)
		},
	}
	return []AuthMethodHandler{nickHandler}
}
