package metools

import (
	"testing"
	"time"
)

type User struct {
	ID        int64     `map:"id" json:"id"`
	Username  string    `json:"user_name"`
	Password  string    `json:"password"`
	Logintime time.Time `json:"-"`
	Fields    []int     `json:"fields"`
	A         A
}

type A struct {
	B string
}

func TestStruct2Map(t *testing.T) {
	u := []User{
		User{
			ID:        1,
			Username:  "tom",
			Password:  "12dbau7",
			Logintime: time.Now(),
			Fields:    []int{1, 2, 4},
			A: A{
				B: "its b",
			},
		},
	}

	m := Struct2Map(&u)
	t.Logf("%+v", m)
}

func TestStructMarshalMap(t *testing.T) {
	u := []User{
		User{
			ID:        1,
			Username:  "tom",
			Password:  "12dbau7",
			Logintime: time.Now(),
			Fields:    []int{1, 2, 4},
			A: A{
				B: "its b",
			},
		},
	}

	m := StructMarshaalToMap(&u)
	t.Logf("%+v", m)
}
