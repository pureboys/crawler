package main

import (
	"my_pritice/crawler/engine"
	"my_pritice/crawler/scheduler"
	"my_pritice/crawler/zhenai/parser"
	"my_pritice/crawler_distributed/persist/client"
	"fmt"
	"my_pritice/crawler_distributed/config"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
