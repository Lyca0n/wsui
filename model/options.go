package model

type Options struct {
	OriginHeader string `json:"Origin-Header"`
	SendAs       string `json:"Send-Content-As"`
	ConsumeAs    string `json:"Get-Content-As"`
	Reconnect    bool   `json:"AutoReconnect"`
}

func DefaultOpts() *Options {
	var ops = &Options{
		Reconnect: true,
	}
	return ops
}
