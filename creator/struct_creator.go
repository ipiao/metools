package creator

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"time"

	metools "github.com/ipiao/metools/utils"
)

var defaultTmpl, _ = template.ParseFiles("templ/struct.tpl")

// common types
var (
	TypeString = reflect.TypeOf("")
	TypeInt    = reflect.TypeOf(0)
	TypeBool   = reflect.TypeOf(false)
	TypeTime   = reflect.TypeOf(time.Now())
)

// StructCreator create struct
type StructCreator struct {
	Imports     []string
	Pkg         string
	dir         string
	filename    string
	tmpl        *template.Template
	output      io.Writer
	isouputfile bool
	Struct
}

// Struct describe struct
type Struct struct {
	Name    string
	Fields  []*StructField
	Comment string
}

// StructField describe field
type StructField struct {
	Name  string
	Annoy bool
	Type  string
	Tags  []StrcutTag
	t     reflect.Type
}

// StrcutTag is tag
type StrcutTag struct {
	Name  string
	Value string
}

// NewStruct for struct
// pkg is package in file
// name is name of struct
func NewStruct(pkg, name string) *StructCreator {
	if pkg == "" || name == "" {
		panic("pkg or name can not ne empty")
	}
	var filename = metools.SnakeName(name)
	var commont = fmt.Sprintf("is %s", filename)
	filename = filename + ".go"

	return &StructCreator{
		dir:      pkg,
		filename: filename,
		Pkg:      pkg,
		tmpl:     defaultTmpl,
		Struct: Struct{
			Name:    name,
			Comment: commont,
		},
	}
}

// SetTmpl set template
func (s *StructCreator) SetTmpl(tmpl *template.Template) {
	s.tmpl = tmpl
}

// SetPath set dir and filename
func (s *StructCreator) SetPath(p string) {
	if p == "" {
		panic("path can not ne empty")
	}
	dir := filepath.Dir(p)
	fname := filepath.Base(p)
	if !strings.HasSuffix(fname, ".go") {
		fname = fname + ".go"
	}
	s.dir = dir
	s.filename = fname
}

// GetImports get imports
func (s *StructCreator) GetImports() []string {
	if len(s.Imports) == 0 {
		for i := range s.Fields {
			if s.Fields[i].t.PkgPath() != "" && metools.Index(s.Fields[i].t.PkgPath(), s.Imports) == -1 {
				s.Imports = append(s.Imports, s.Fields[i].t.PkgPath())
			}
		}
	}
	return s.Imports
}

// Exec impl Creator
func (s *StructCreator) Exec() error {
	s.GetImports()

	if s.output == nil {
		path := filepath.Join(s.dir, s.filename)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err := os.MkdirAll(s.dir, os.ModePerm)
			if err != nil {
				return err
			}
			output, err := os.Create(path)
			if err != nil {
				return err
			}
			s.SetOutput(output, true)
		} else {
			output, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				return err
			}
			s.SetOutput(output, true)
		}
	}

	err := s.tmpl.Execute(s.output, s)
	if err != nil {
		return err
	}
	if s.isouputfile {
		return s.format()
	}
	return nil
}

// SetOutput set output
func (s *StructCreator) format() error {
	cmd := exec.Command("gofmt", "-w", filepath.Join(s.dir, s.filename))
	err := cmd.Run()
	return err
}

// SetOutput set output
func (s *StructCreator) SetOutput(output io.Writer, isfile bool) {
	s.output = output
	s.isouputfile = isfile
}

// AddField set template
func (s *Struct) AddField(fields ...*StructField) {
	for _, field := range fields {
		if field.t == nil {
			panic("filed's type is undefined")
		}
		for _, f := range s.Fields {
			if f.Name == field.Name {
				panic(fmt.Sprintf("name %s has exists instruct", field.Name))
			}
		}
	}
	s.Fields = append(s.Fields, fields...)
}

// NewStructField construct a field
func NewStructField(name string, t reflect.Type) *StructField {
	if t == nil {
		panic("t in field cannot be nil")
	}
	var f = &StructField{
		Name: name,
		t:    t,
	}
	if name == "" {
		f.Annoy = true
	}
	f.Type = t.String()
	return f
}

// AddTag for struct
func (s *StructField) AddTag(name, value string) *StructField {
	s.Tags = append(s.Tags, StrcutTag{
		Name:  name,
		Value: value,
	})
	return s
}
