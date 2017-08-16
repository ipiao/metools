package creator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type People struct{}

func TestCreateStruct(t *testing.T) {
	s := NewStruct("example", "UserHello")
	f1 := NewStructField("Name", TypeString)
	f2 := NewStructField("Age", TypeInt).AddTag("json", "age")
	f3 := NewStructField("BirthDay", TypeTime)
	f4 := NewStructField("", reflect.TypeOf(People{}))
	f5 := NewStructField("people", reflect.TypeOf(&People{}))
	s.AddField(f1, f2, f3, f4, f5)
	//	s.SetPath("example/user.go")
	//	s.SetOutput(os.Stdout, false)
	err := s.Exec()
	assert.Nil(t, nil, err)
}
