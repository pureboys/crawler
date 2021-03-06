package main

import (
	"my_pritice/crawler/engine"
	"my_pritice/crawler/scheduler"
	"my_pritice/crawler/zhenai/parser"
	"my_pritice/crawler/persist"
	"my_pritice/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
