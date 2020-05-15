package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"gopkg.in/yaml.v2"
)

func main() {
	bytes, err := ioutil.ReadFile("/home/ellie/Code/upenn/user-dataset-upload-service/schema/library.raml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	foo := raml.Library{}
	if err := yaml.Unmarshal(bytes, &foo); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	enc := yaml.NewEncoder(os.Stdout)
	enc.Encode(foo)
	fmt.Println(foo.Types["AnotherType"])
}
