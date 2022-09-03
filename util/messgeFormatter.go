package util

import (
	"bytes"
	"encoding/json"
)

func FormatJson(msg []byte) string {
	var formatted bytes.Buffer
	if err := json.Indent(&formatted, msg, "", "    "); err != nil {
		return ""
	}
	return string(formatted.Bytes())
}

func FormatDefault(msg []byte) string {
	var formatted []byte
	sets := len(msg) / 90
	i := 0
	for ; i <= sets; i = i + 90 {
		formatted = append(formatted, msg[i*90:i+90]...)
		formatted = append(formatted, '\n')
	}

	formatted = append(formatted, msg[i:]...)
	return string(formatted)
}
