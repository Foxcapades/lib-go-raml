package main

import (
	"github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml/rparse"
	"os"
)

func main() {
	res, err := simple.GetRequest("https://raw.githubusercontent.com/raml-org/raml-examples/master/annotations/advanced.raml").
		Submit().
		GetRawResponse()
	check(err)
	defer res.Body.Close()

	api, err := rparse.ParseApiDoc(res.Body)
	check(err)

	check(api.WriteRAML(os.Stdout))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
