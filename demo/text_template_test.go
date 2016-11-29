package demo

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

func HTMLEscaoeString(s string) {
	var res = template.HTMLEscapeString(s)
	fmt.Println(res)
}

func ParseFiles() {

	muban1 := "woshim1"
	muban2 := "woshim2"

	temp, err := template.ParseFiles("../file/test_template1.tpl", "../file/test_template2.tpl")
	if err != nil {
		panic(err)
	}
	//temp.Execute()
	temp, err = temp.New("M1").Parse(muban1)
	if err != nil {
		panic(err)
	}
	temp, err = temp.New("M2").Parse(muban2)
	if err != nil {
		panic(err)
	}
	err = temp.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	//log.Println(err)
	//log.Println(temp.Name())
}

func TestTemplate(t *testing.T) {
	//HTMLEscaoeString("<label>hello<label/>")
	ParseFiles()
}
