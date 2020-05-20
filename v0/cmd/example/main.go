package main

import (
	"fmt"
	"github.com/Foxcapades/Go-ChainRequest/simple"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rparse"
)

func main() {
	res, err := simple.GetRequest("https://raw.githubusercontent.com/raml-org/raml-examples/master/others/tutorial-jukebox-api/jukebox-api.raml").
		Submit().GetRawResponse()

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	api, err := rparse.ParseApiDoc(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(api.Traits())
}
