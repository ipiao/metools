package mutils

import (
	"testing"
	"time"
)

type User struct {
	ID        int64  `map:"id" json:"id"`
	Username  string `json:"user_name"`
	password  string
	Logintime time.Time `json:"logintime"`
	Fields    []int     `json:"fields"`
	A         A         `json:"a"`
}

type A struct {
	B string `json:"b"`
}

func TestStruct2Map(t *testing.T) {

	u := User{
		ID:        1,
		Username:  "tom",
		password:  "12dbau7",
		Logintime: time.Now(),
		Fields:    []int{1, 2, 4},
		A: A{
			B: "its b",
		},
	}
	m := Struct2Map(&u)
	t.Logf("%+v", m)
}

func BenchmarkStruct2Map(b *testing.B) {
	u := User{
		ID:        1,
		Username:  "tom",
		password:  "12dbau7",
		Logintime: time.Now(),
		Fields:    []int{1, 2, 4},
		A: A{
			B: "its b",
		},
	}
	for i := 0; i < b.N; i++ {
		Struct2Map(&u)
	}
}

func BenchmarkStructMarshalMap(b *testing.B) {
	u := User{
		ID:        1,
		Username:  "tom",
		password:  "12dbau7",
		Logintime: time.Now(),
		Fields:    []int{1, 2, 4},
		A: A{
			B: "its b",
		},
	}
	for i := 0; i < b.N; i++ {
		StructMarshalToMap(&u)
	}
}

func TestStructMarshalMap(t *testing.T) {

	u := User{
		ID:        1,
		Username:  "tom",
		password:  "12dbau7",
		Logintime: time.Now(),
		Fields:    []int{1, 2, 4},
		A: A{
			B: "its b",
		},
	}

	m := StructMarshalToMap(&u)
	t.Logf("%+v", m)
}

func BenchmarkTransFieldName(b *testing.B) {
	var name = "NsadaKsadaNadasMss"
	for i := 0; i < b.N; i++ {
		TransFieldName(name)
	}
}

func BenchmarkSnakeName(b *testing.B) {
	var name = "NsadaKsadaNadasMss"
	for i := 0; i < b.N; i++ {
		SnakeName(name)
	}
}

// func TestBase64(t *testing.T) {
// 	res := Base64Encode("ykk@#1001")
// 	t.Log(res)
// }

func TestRemoveDuplicatesInts(t *testing.T) {
	arr := []int{1, 1, 2, 3, 4, 4}
	arr = RemoveDuplicatesInts(arr)
	t.Log(arr)
}
