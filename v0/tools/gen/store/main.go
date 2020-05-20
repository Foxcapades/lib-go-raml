package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

const (
	iFilePrefix = "v0/pkg/raml/x-gen-map-"
	mFilePrefix = "v0/internal/raml/x-gen-map-"
	fileSuffix  = ".go"
)

var now = time.Now().Format(time.RFC3339Nano)

func ramlRow(Name, KType, KName string) row {
	return row{
		Time:  now,
		Name:  Name,
		Type:  Name,
		Kind:  Name,
		KType: KType,
		KName: KName,
		Raml:  true,
	}
}

func nonRow(Name, Type, Kind, KType, KName string) row {
	return row{
		Time:  now,
		Name:  Name,
		Type:  Type,
		KType: KType,
		KName: KName,
		Kind:  Kind,
	}
}

var types = []row{
	ramlRow("Annotation", "string", "String"),
	ramlRow("DataType", "string", "String"),
	ramlRow("ArrayExample", "string", "String"),
	ramlRow("BoolExample", "string", "String"),
	ramlRow("CustomExample", "string", "String"),
	ramlRow("DateOnlyExample", "string", "String"),
	ramlRow("DatetimeExample", "string", "String"),
	ramlRow("DatetimeOnlyExample", "string", "String"),
	ramlRow("FileExample", "string", "String"),
	ramlRow("IntegerExample", "string", "String"),
	ramlRow("NumberExample", "string", "String"),
	ramlRow("ObjectExample", "string", "String"),
	ramlRow("StringExample", "string", "String"),
	ramlRow("TimeOnlyExample", "string", "String"),
	ramlRow("UnionExample", "string", "String"),
	ramlRow("Facet", "string", "String"),
	nonRow("Untyped", "interface{}", "Untyped", "string", "String"),
	ramlRow("Property", "string", "String"),
	nonRow("String", "string", "String", "string", "String"),
	nonRow("Any", "interface{}", "Untyped", "interface{}", "Untyped"),
}

func main() {
	iTpl := template.Must(template.ParseFiles("v0/gen/store/interface.tpl"))
	mTpl := template.Must(template.ParseFiles("v0/gen/store/implementation.tpl"))

	for _, row := range types {
		lName := strings.ToLower(row.Name)
		iFile, err := os.Create(iFilePrefix + lName + fileSuffix)
		mFile, err := os.Create(mFilePrefix + lName + fileSuffix)

		check(err)

		check(iTpl.ExecuteTemplate(iFile, "map", row))
		check(mTpl.ExecuteTemplate(mFile, "impl", row))

		iFile.Close()
		mFile.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type row struct {
	Time  string
	Type  string
	Name  string
	TName string
	KType string
	KName string
	Raml  bool
	Kind  string
}
