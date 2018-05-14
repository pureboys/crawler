package client

import (
	"my_pritice/crawler/engine"
	"my_pritice/crawler_distributed/rpcsupport"
	"fmt"
	"my_pritice/crawler_distributed/config"
	"my_pritice/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
