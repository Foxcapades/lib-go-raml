//+build ignore

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
	{Name: "Array", Type: "[]interface{}", Time: now},
	{Name: "Bool", Type: "bool", Time: now},
	{Name: "Custom", Type: "interface{}", Time: now},
	{Name: "DateOnly", Type: "string", Time: now},
	{Name: "TimeOnly", Type: "string", Time: now},
	{Name: "DatetimeOnly", Type: "string", Time: now},
	{Name: "Datetime", Type: "string", Time: now},
	{Name: "File", Type: "interface{}", Time: now},
	{Name: "Integer", Type: "int64", Time: now},
	{Name: "Number", Type: "float64", Time: now},
	{Name: "Object", Type: "interface{}", Time: now},
	{Name: "String", Type: "string", Time: now},
	{Name: "Union", Type: "interface{}", Time: now},
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
	Name string
	Type string
	Time string
}
