package main

import (
	"fmt"
	"regexp"
)

const text = `My email is linbin@smzdm.com
email1 abc@dfc.org
email2 ddd@qq.com
email3 dd@qq.com.cn
`

func main() {
	//re, err := regexp.Compile("linbin@smzmd.com")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
