package main

import (
	"fmt"
	"my_pritice/crawler/engine"
	"my_pritice/crawler/scheduler"
	"my_pritice/crawler/zhenai/parser"
	"my_pritice/crawler_distributed/config"
	itemsaver "my_pritice/crawler_distributed/persist/client"
	worker "my_pritice/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
