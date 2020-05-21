package main

import (
	"os"
	"text/template"
)

const (
	version   = "v0"
	genPath   = version + "/gen"
	docPath   = genPath + "/doc/"
	traitPath = genPath + "/traits-i/"
	iPath     = version + "/pkg/raml/"
)

var docFiles = [][3]string{
	{"api-i.tpl", "api-interface", "APISpec"},
	{"documentation-i.tpl", "documentation-interface", "Documentation"},
	{"library-i.tpl", "library-interface", "Library"},
	{"method-i.tpl", "method-interface", "Method"},
	{"overlay-i.tpl", "overlay-interface", "Overlay"},
	{"response-i.tpl", "response-interface", "Response"},
}

func main() {
	files := make([]string, 0, len(docFiles))

	for _, v := range docFiles {
		files = append(files, docPath+v[0])
	}

	temp := template.Must(template.Must(template.ParseFiles(files...)).
		ParseGlob(traitPath + "*"))

	for _, v := range docFiles {
		file, err := os.Create(iPath + v[2] + ".go")
		check(err)
		check(temp.ExecuteTemplate(file, v[1], nil))
		check(file.Close())
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
