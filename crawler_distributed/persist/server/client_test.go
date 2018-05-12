package main

import (
	"testing"
	"my_pritice/crawler_distributed/rpcsupport"
	"my_pritice/crawler/engine"
	"my_pritice/crawler/model"
	"time"
	"my_pritice/crawler_distributed/config"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1344440064",
		Type: "zhenai",
		Id:   "1344440064",
		Payload: model.Profile{
			Age:        39,
			Height:     173,
			Weight:     65,
			Income:     "5001-8000元",
			Gender:     "男",
			Name:       "真心找对象",
			Xinzuo:     "处女座",
			Occupation: "通信技术",
			Marriage:   "离异",
			House:      "打算婚后购房",
			Hokou:      "安徽阜阳",
			Education:  "高中及以下",
			Car:        "已购车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, err: %s", result, err)
	}

}
