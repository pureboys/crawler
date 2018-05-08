package parser

import (
	"testing"
	"io/ioutil"
	"my_pritice/crawler/model"
	"my_pritice/crawler/engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/1344440064", "真心找对象")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]

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

	if actual != expected {
		t.Errorf("expected %v but was %v", expected, actual)
	}

}
