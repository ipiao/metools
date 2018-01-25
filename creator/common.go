package creator

import (
	"path/filepath"
	"reflect"
	"runtime"
	"text/template"
	"time"
)

var (
	_, f, _, _     = runtime.Caller(0)
	_structTplFile = filepath.Join(filepath.Dir(f), "tmpl/struct.tpl")
	_funcTplFile   = filepath.Join(filepath.Dir(f), "tmpl/func.tpl")
)

var defaultFuncTmpl, _ = template.ParseFiles(_funcTplFile)
var defaultStructTmpl, _ = template.ParseFiles(_structTplFile)

// common types
var (
	TypeString = reflect.TypeOf("")
	TypeInt    = reflect.TypeOf(0)
	TypeBool   = reflect.TypeOf(false)
	TypeTime   = reflect.TypeOf(time.Now())
)
