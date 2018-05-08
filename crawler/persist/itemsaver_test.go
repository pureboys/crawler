package persist

import (
	"testing"
	"my_pritice/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"my_pritice/crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}
