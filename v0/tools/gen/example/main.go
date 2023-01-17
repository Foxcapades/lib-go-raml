//go:build ignore
// +build ignore

package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

const (
	iFilePrefix = "v0/pkg/raml/x-gen-ex-"
	mFilePrefix = "v0/internal/raml/x-gen-ex-"
	fileSuffix  = ".go"
)

var now = time.Now().Format(time.RFC3339Nano)

var types = []exampleProps{
	{Name: "Any", Type: "interface{}", EType: "Untyped", Time: now, UseOption: true},
	{Name: "Array", Type: "[]interface{}", EType: "[]interface{}", Time: now, UseOption: false},
	{Name: "Bool", Type: "bool", EType: "Bool", Time: now, UseOption: true},
	{Name: "Custom", Type: "interface{}", EType: "Untyped", Time: now, UseOption: true},
	{Name: "DateOnly", Type: "string", EType: "String", Time: now, UseOption: true},
	{Name: "TimeOnly", Type: "string", EType: "String", Time: now, UseOption: true},
	{Name: "DatetimeOnly", Type: "string", EType: "String", Time: now, UseOption: true},
	{Name: "Datetime", Type: "string", EType: "String", Time: now, UseOption: true},
	{Name: "File", Type: "interface{}", EType: "Untyped", Time: now, UseOption: true},
	{Name: "Integer", Type: "int64", EType: "Int64", Time: now, UseOption: true},
	{Name: "Number", Type: "float64", EType: "Float64", Time: now, UseOption: true},
	{Name: "Object", Type: "interface{}", EType: "Untyped", Time: now, UseOption: true},
	{Name: "String", Type: "string", EType: "String", Time: now, UseOption: true},
	{Name: "Union", Type: "interface{}", EType: "Untyped", Time: now, UseOption: true},
}

func main() {
	iTpl := template.Must(template.ParseFiles("v0/gen/example-i.tpl"))
	mTpl := template.Must(template.ParseFiles("v0/gen/example-m.tpl"))
	for i := range types {
		fName := strings.ToLower(types[i].Name)

		iFile, err := os.Create(iFilePrefix + fName + fileSuffix)
		check(err)
		mFile, err := os.Create(mFilePrefix + fName + fileSuffix)
		check(err)

		check(iTpl.ExecuteTemplate(iFile, "example", types[i]))
		check(mTpl.ExecuteTemplate(mFile, "example", types[i]))
		iFile.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type exampleProps struct {
	Name      string
	Type      string
	EType     string
	Time      string
	UseOption bool
}
