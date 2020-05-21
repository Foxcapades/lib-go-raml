package main

import (
	"github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rparse"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	res, err := simple.GetRequest("https://raw.githubusercontent.com/raml-org/raml-examples/master/others/tutorial-jukebox-api/jukebox-api.raml").
		Submit().
		GetRawResponse()
	check(err)
	defer res.Body.Close()

	api, err := rparse.ParseApiDoc(res.Body)
	check(err)

	check(yaml.NewEncoder(os.Stdout).Encode(api.Types()))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
