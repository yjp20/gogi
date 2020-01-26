package gogi

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthMethodHandler struct {
	Method  string
	Slug    string
	Handler httprouter.Handle
}

type AuthMethod interface {
	Init(*Context) error
	Template() *template.Template
	Handlers(*Context) []AuthMethodHandler
}

func WithAuthMethod(am AuthMethod) Option {
	return func(c *Context) {
		err := am.Init(c)
		if err != nil {
			log.Fatal(err)
		}
		c.AuthMethods = append(c.AuthMethods, am)
	}
}

var guestAuthTemplate = `<div class="card-content">
	<form class="form" method="post" action="{{.Context.Prefix}}">
		<div class="field">
			<label class="label"> Nickname </label>
			<div class="control">
				<input name="nick" class="input" type="text" placeholder="Guest Login">
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
			}

		},
	}
	return []AuthMethodHandler{nickHandler}
}

var loginAuthTemplate = `<div class="card-content">
	<div class="tabs is-centered">
		<ul style="border-bottom: 1px solid #ccc">
			<li><a onclick="document.getElementById('loginForm').style = ''; document.getElementById('signupForm').style = 'display: none';">Login</a></li>
			<li><a onclick="document.getElementById('loginForm').style = 'display: none'; document.getElementById('signupForm').style = '';">Signup</a></li>
		</ul>
	</div>
	<div id="loginSignupForms">
		<div id="loginForm">
			<form class="form">
				<div class="field">
					<label class="label"> Username </label>
					<div class="control">
						<input name="username" class="input" type="text" placeholder="Username">
					</div>
				</div>
				<div class="field">
					<label class="label"> Password </label>
					<div class="control">
						<input name="password" class="input" type="password" placeholder="Password">
					</div>
				</div>
				<div class="field">
					<div class="control">
						<button class="button is-fullwidth"> Login </button>
					</div>
				</div>
			</form>
		</div>
		<div id="signupForm" style="display: none">
			<form class="form">
				<div class="field">
					<label class="label"> Username </label>
					<div class="control">
						<input name="username" class="input" type="text" placeholder="Username">
					</div>
				</div>
				<div class="field">
					<label class="label"> Email </label>
					<div class="control">
						<input name="email" class="input" type="email" placeholder="email">
					</div>
				</div>
				<div class="field">
					<label class="label"> Password </label>
					<div class="control">
						<input name="password" class="input" type="passowrd" placeholder="Password">
					</div>
				</div>
				<div class="field">
					<div class="control">
						<button class="button is-fullwidth"> Signup </button>
					</div>
				</div>
			</form>
		</div>
	</div>
</div>`

type LoginAuth struct {
	TemplateParsed *template.Template
}

func (a *LoginAuth) Init(c *Context) error {
	var err error
	if c.DB == nil {
		log.Fatal("Login Auth requires a DB provider")
	}
	if a.TemplateParsed == nil {
		a.TemplateParsed = template.New("")
		_, err = a.TemplateParsed.Parse(loginAuthTemplate)
	}
	return err
}

func (a *LoginAuth) Template() *template.Template {
	return a.TemplateParsed
}

func (a *LoginAuth) Handlers(c *Context) []AuthMethodHandler {
	nickHandler := AuthMethodHandler{
		"POST", "", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		},
	}
	return []AuthMethodHandler{nickHandler}
}
