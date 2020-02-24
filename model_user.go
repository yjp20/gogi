package gogi

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetID() uint
	GetName() string
	SetName(string)
	GetNick() string
	SetNick(string)
	GetTemp() bool
	SetTemp(bool)
	GetEmail() string
	SetEmail(string)
	GetShortID() string
	SetShortID()
	IsPassword(string) bool
	SetPassword(string)
}

type UserModel struct {
	gorm.Model

	Name     string
	Nick     string
	Temp     bool
	Email    string
	ShortID  string
	Password []byte
}

func (u *UserModel) GetID() uint {
	return u.ID
}

func (u *UserModel) GetName() string {
	return u.Name
}

func (u *UserModel) SetName(s string) {
	u.Name = s
}

func (u *UserModel) GetNick() string {
	return u.Nick
}

func (u *UserModel) SetNick(s string) {
	u.Nick = s
}

func (u *UserModel) GetTemp() bool {
	return u.Temp
}

func (u *UserModel) SetTemp(b bool) {
	u.Temp = b
}

func (u *UserModel) GetShortID() string {
	return u.ShortID
}

func (u *UserModel) SetShortID() {
	u.ShortID, _ = shortid.Generate()
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
