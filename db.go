package gogi

import (
	"log"

	"github.com/jinzhu/gorm"
)

func WithDBProvider(db, connstr string) Option {
	return func(c *Context) {
		var err error
		c.DB, err = gorm.Open(db, connstr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
