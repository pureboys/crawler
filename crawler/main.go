package main

import (
	"my_pritice/crawler/engine"
	"my_pritice/crawler/scheduler"
	"my_pritice/crawler/zhenai/parser"
	"my_pritice/crawler/persist"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
