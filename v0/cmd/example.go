package main

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/pkg/raml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("/home/ellie/Code/upenn/user-dataset-upload-service/schema/url/user-dataset/post-response.raml")
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
	fmt.Println(foo)
}
