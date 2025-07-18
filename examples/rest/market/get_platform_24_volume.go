package main

import (
	"log"

	"github.com/bryanchen463/go-okx/examples/rest"
	"github.com/bryanchen463/go-okx/rest/api/market"
)

func main() {
	req, resp := market.NewGetPlatform24Volume()
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetPlatform24VolumeResponse))
}
