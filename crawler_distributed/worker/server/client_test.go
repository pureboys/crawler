package main

import (
	"testing"
	"my_pritice/crawler_distributed/rpcsupport"
	"my_pritice/crawler_distributed/worker"
	"time"
	"my_pritice/crawler_distributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1675198875",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "时光隧道",
		},
	}

	var result worker.ParseResult

	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
