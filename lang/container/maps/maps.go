package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println("m,m2,m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["cource"]=`, courseName)

}
