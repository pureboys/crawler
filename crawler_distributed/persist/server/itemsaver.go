package main

import (
	"my_pritice/crawler_distributed/rpcsupport"
	"my_pritice/crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
	"my_pritice/crawler_distributed/config"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
