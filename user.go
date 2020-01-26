package gogi

import (
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetName() string
	SetName(string)
	GetTemp() bool
	SetTemp(bool)
	GetEmail() string
	SetEmail(string)
	IsPassword(string) bool
	SetPassword(string)
}

type UserModel struct {
	gorm.Model

	Name     string
	Temp     bool
	Email    string
	Password []byte
}

func (u *UserModel) GetName() string {
	return u.Name
}

func (u *UserModel) SetName(s string) {
	u.Name = s
}

func (u *UserModel) GetTemp() bool {
	return u.Temp
}

func (u *UserModel) SetTemp(b bool) {
	u.Temp = b
}

func (u *UserModel) GetEmail() string {
	return u.Email
}

func (u *UserModel) SetEmail(email string) {
	u.Email = email
}

func (u *UserModel) IsPassword(s string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(s))
	if err == nil {
		return true
	}
	return false
}

func (u *UserModel) SetPassword(s string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = hash
}

func WithUserModel(u User) Option {
	return func(c *Context) {
		c.UserModel = u
	}
}
