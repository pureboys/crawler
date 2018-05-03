package parser

import (
	"my_pritice/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])

		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
		//fmt.Printf("City: %s, URL: %s \n", m[2], m[1])
	}
	return result
}
