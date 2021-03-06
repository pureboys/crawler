package parser

import (
	"my_pritice/crawler/engine"
	"regexp"
	"my_pritice/crawler_distributed/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
	}
	return result
}
