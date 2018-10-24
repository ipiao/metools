package example

import (
	"time"

	"github.com/ipiao/metools/creator"
)

// User is user
type User struct {
	Name     string
	Age      int `json:"age"`
	BirthDay time.Time
	creator.People
	people *creator.People
}

func (u User) Hello(s string) string {
	return "hello"
}
