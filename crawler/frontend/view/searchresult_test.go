package view

import (
	"testing"
	"my_pritice/crawler/frontend/model"
	common "my_pritice/crawler/model"
	"os"
	"my_pritice/crawler/engine"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	file, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1344440064",
		Type: "zhenai",
		Id:   "1344440064",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(file, page)
	if err != nil {
		panic(err)
	}
}
