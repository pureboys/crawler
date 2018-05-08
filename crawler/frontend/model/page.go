package model

import "my_pritice/crawler/engine"

type SearchResult struct {
	Hits int
	Start int
	Items []engine.Item
}

