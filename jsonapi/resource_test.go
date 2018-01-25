package jsonapi

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	ID   int
	Name string
}

func (User) Type() string {
	return "user"
}

type User2 struct {
	Id string
}

type User3 struct {
	User2
}

type User4 struct {
}

type User5 struct {
	User4
}

func (User2) Type() string {
	return "User"
}

type Interface struct{}

func TestGetType(t *testing.T) {

	assert.Equal(t, "user", GetType(User{}))
	assert.Equal(t, "user", GetType(&User{}))
	assert.Equal(t, "User", GetType(&User2{}))
	assert.Equal(t, "User", GetType((&User3{})))
	assert.Equal(t, "user5", GetType((&User5{})))
	assert.Equal(t, "user", GetType([]User{User{}}))
	assert.Equal(t, "user", GetType([]*User{&User{}}))

	assert.Equal(t, "interface", GetType(interface{}(Interface{})))
	assert.Equal(t, "interface", GetType(interface{}(&Interface{})))
	assert.Equal(t, "User", GetType(interface{}(&User3{})))

	assert.Equal(t, "", GetType([]interface{}{&Interface{}}))
	assert.Equal(t, "", GetType(0))
	assert.Equal(t, "", GetType(false))
	assert.Equal(t, "", GetType(""))
	assert.Equal(t, "", GetType(interface{}(0)))

}

func BenchmarkGetType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assert.Equal(b, "user", GetType(User{}))
		assert.Equal(b, "user", GetType(&User{}))
		assert.Equal(b, "User", GetType(&User2{}))
		assert.Equal(b, "user", GetType([]User{User{}}))
		assert.Equal(b, "user", GetType([]*User{&User{}}))

		assert.Equal(b, "interface", GetType(interface{}(Interface{})))
		assert.Equal(b, "interface", GetType(interface{}(&Interface{})))

		assert.Equal(b, "", GetType([]interface{}{&Interface{}}))
		assert.Equal(b, "", GetType(0))
		assert.Equal(b, "", GetType(false))
		assert.Equal(b, "", GetType(""))
		assert.Equal(b, "", GetType(interface{}(0)))
	}
}

func TestSyncGetType(t *testing.T) {
	for i := 0; i < 20; i++ {
		assert.Equal(t, "user", GetType(User{}))
		assert.Equal(t, "user", GetType(&User{}))
		assert.Equal(t, "User", GetType(&User2{}))
		assert.Equal(t, "user", GetType([]User{User{}}))
		assert.Equal(t, "user", GetType([]*User{&User{}}))

		assert.Equal(t, "interface", GetType(interface{}(Interface{})))
		assert.Equal(t, "interface", GetType(interface{}(&Interface{})))

		assert.Equal(t, "", GetType([]interface{}{&Interface{}}))
		assert.Equal(t, "", GetType(0))
		assert.Equal(t, "", GetType(false))
		assert.Equal(t, "", GetType(""))
		assert.Equal(t, "", GetType(interface{}(0)))
	}
	assert.Equal(t, len(typeMap), 7)
	assert.Equal(t, typeMap[reflect.TypeOf(User{})], "user")
}

func TestTransModelToData(t *testing.T) {
	t.Log(transModelToData(User{ID: 212312, Name: "tom"}))
	t.Log(transModelToData(&User{ID: 212312, Name: "tom"}))
	t.Log(transModelToData([]User{User{ID: 212312, Name: "tom"}}))
	t.Log(transModelToData([]*User{&User{ID: 212312, Name: "tom"}}))
	t.Log(transModelToData(interface{}(&User{ID: 212312, Name: "tom"})))
	t.Log(transModelToData(interface{}(User{ID: 212312, Name: "tom"})))
	t.Log(transModelToData(map[string]User{"a": User{ID: 212312, Name: "tom"}}))
	t.Log(transModelToData(User2{Id: "212312"}))
	t.Log(transModelToData(User5{}))

	t.Log(transModelToData(map[string]interface{}{"a": User{ID: 212312, Name: "tom"}}))
	t.Log(idIndexMap)
	t.Log(typeMap)
}
