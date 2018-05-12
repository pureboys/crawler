package client

import (
	"my_pritice/crawler/engine"
	"log"
	"my_pritice/crawler_distributed/rpcsupport"
	"my_pritice/crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: GET item #%d:%v", itemCount, item)
			itemCount++

			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: err saving item %v, %v", item, err)
			}
		}
	}()
	return out, nil
}
