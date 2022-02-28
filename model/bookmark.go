package model

import (
	"net/url"
)

type Bookmark struct {
	Url  url.URL
	Name string
}

func (b *Bookmark) String() string {
	return b.Name + " @ " + b.Url.String()
}
