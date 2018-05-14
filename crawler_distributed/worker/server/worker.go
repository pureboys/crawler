package main

import (
	"my_pritice/crawler_distributed/rpcsupport"
	"fmt"
	"my_pritice/crawler_distributed/config"
	"my_pritice/crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}),
	)
}
