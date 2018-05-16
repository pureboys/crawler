package client

import (
	"my_pritice/crawler/engine"
	"my_pritice/crawler_distributed/config"
	"my_pritice/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(request engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(request)

		var sResult worker.ParseResult
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
