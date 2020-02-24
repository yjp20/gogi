package gogi

import (
	"crypto/rand"
	"log"
	"net/http"
	"reflect"

	"github.com/NYTimes/gziphandler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Context struct {
	Name        string
	AuthMethods []AuthMethod
	Middlewares []func(http.Handler) http.Handler
	Prefix      string
	Description string
	DB          *gorm.DB
	Rooms       map[string]Room
	Version     string
	Secret      []byte

	UserModel    User
	RoomModel    Room
	ManagerModel Manager
}

func (c *Context) Init() error {
	c.Rooms = make(map[string]Room)

	if c.DB == nil {
		var err error
		c.DB, err = gorm.Open("sqlite3", ":memory:")
		if err != nil {
			return err
		}
	}

	if len(c.Secret) == 0 {
		_, err := rand.Read(c.Secret)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Context) NewUser() User {
	a := reflect.ValueOf(c.UserModel).Elem() // Gets the user supplied model
	u := reflect.New(a.Type()).Interface().(User)
	return u
}

func (c *Context) NewRoom() Room {
	a := reflect.ValueOf(c.RoomModel).Elem() // Gets the user supplied model
	r := reflect.New(a.Type()).Interface().(Room)
	return r
}

type Option func(*Context)

func WithDBProvider(db, connstr string) Option {
	return func(c *Context) {
		var err error
		c.DB, err = gorm.Open(db, connstr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func WithGzip() Option {
	return func(c *Context) {
		c.Middlewares = append(c.Middlewares, gziphandler.GzipHandler)
	}
}

func WithPrefix(prefix string) Option {
	return func(c *Context) {
		c.Prefix = prefix
	}
}

func WithDescription(desc string) Option {
	return func(c *Context) {
		c.Description = desc
	}
}

func WithVersion(version string) Option {
	return func(c *Context) {
		c.Version = version
	}
}

func WithAuthMethod(am AuthMethod) Option {
	return func(c *Context) {
		c.AuthMethods = append(c.AuthMethods, am)
	}
}
