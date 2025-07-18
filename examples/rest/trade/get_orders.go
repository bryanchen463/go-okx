package main

import (
	"log"

	"github.com/bryanchen463/go-okx/examples/rest"
	"github.com/bryanchen463/go-okx/rest/api/trade"
)

func main() {
	param := &trade.GetOrderParam{
		InstId: "OKB-USDT",
		OrdId:  "501163171776954368",
	}
	req, resp := trade.NewGetOrder(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*trade.GetOrderResponse))
}
