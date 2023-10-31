package model

import "sort"

type Options struct {
	SendAs    ContentOptions `json:"Send-Content-As"`
	ConsumeAs ContentOptions `json:"Get-Content-As"`
}

func DefaultOpts() *Options {
	var ops = &Options{
		ConsumeAs: JSON,
		SendAs:    JSON,
	}
	return ops
}

type ContentOptions int8

const (
	JSON ContentOptions = iota
	Text
)

func (c ContentOptions) String() string {
	return [...]string{"Text", "JSON"}[c]
}

func ValueFromString(val string) ContentOptions {
	return ContentOptions(sort.StringSlice([]string{"Text", "JSON"}).Search(val))
}
