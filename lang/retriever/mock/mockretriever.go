package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
