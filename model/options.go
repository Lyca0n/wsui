package model

type Options struct {
	OriginHeader string `json:"Origin-Header"`
	SendAs       string `json:"Send-Content-As"`
	ConsumeAs    string `json:"Get-Content-As"`
}

func DefaultOpts() *Options {
	var ops = &Options{}
	return ops
}
