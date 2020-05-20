package main

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rparse"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {

	//logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(new(prefixed.TextFormatter))
	foo, err := rparse.ParseLibraryFile("/home/ellie/Code/upenn/user-dataset-upload-service/schema/library.raml")
	if err != nil {
		panic(err)
	}

	enc := yaml.NewEncoder(os.Stdout)
	enc.Encode(foo)
	//yaml.Marshal(foo)
	//fmt.Println(foo.Types())
}
