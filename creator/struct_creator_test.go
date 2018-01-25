package creator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStruct(t *testing.T) {
	s := NewStruct("example", "User")
	f1 := NewStructField("Name", TypeString)
	f2 := NewStructField("Age", TypeInt).AddTag("json", "age")
	f3 := NewStructField("BirthDay", TypeTime)
	f4 := NewStructField("", reflect.TypeOf(People{}))
	f5 := NewStructField("people", reflect.TypeOf(&People{}))
	s.AddField(f1, f2, f3, f4, f5)

	func1 := NewFunc("Hello").SetReceiver2("User").AddArgIn2("s", "string").AddArgOut2("string")
	func1.FuncBody = `return "hello"`
	s.Funcs = append(s.Funcs, func1)
	t.Log(s)
	// structTmpl := template.Must(template.ParseFiles("tmpl/struct.tpl"))
	// funcTmpl := template.Must(template.ParseFiles("tmpl/struct.tpl", "tmpl/func.tpl"))

	//assert.Nil(t, nil, err)
	// s.SetTmpl(funcTmpl)
	//s.SetPath("example/user.go")

	var err error
	// tmpl, _ := template.ParseFiles("tmpl/struct.tpl", "tmpl/func.tpl")
	//s.SetTmpl(funcTmpl)
	// s.SetOutput(os.Stdout, false)
	err = s.Exec()
	assert.Nil(t, nil, err)
}
