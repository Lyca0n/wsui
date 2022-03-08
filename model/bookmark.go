package model

import (
	"net/url"
	"strings"
)

type Bookmark struct {
	Url  url.URL
	Name string
}

func (b *Bookmark) String() string {
	return b.Name + " @ " + b.Url.String()
}

func FilterByTerm(term string, list []Bookmark) []Bookmark {
	var data []Bookmark
	for _, item := range list {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(term)) {
			data = append(data, item)
		}
	}
	return data
}
