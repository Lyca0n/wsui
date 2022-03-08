package model

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type AppState struct {
	ConnectionList []Bookmark
	Messages       []string
	Connection     websocket.Conn
}

func InitAppState() *AppState {
	var initState = &AppState{}
	initState.ConnectionList = []Bookmark{
		{Name: "Home", Url: url.URL{Scheme: "ws", Host: "localhost", Path: ""}},
		{Name: "Store 0020", Url: url.URL{Scheme: "ws", Host: "192.168.0.120", Path: ""}},
		{Name: "Home", Url: url.URL{Scheme: "ws", Host: "localhost", Path: ""}},
		{Name: "Store 0020", Url: url.URL{Scheme: "ws", Host: "192.168.0.120", Path: ""}},
		{Name: "Home", Url: url.URL{Scheme: "ws", Host: "localhost", Path: ""}},
		{Name: "Store 0020", Url: url.URL{Scheme: "ws", Host: "192.168.0.120", Path: ""}},
		{Name: "Store 0020", Url: url.URL{Scheme: "ws", Host: "192.168.0.120", Path: ""}},
	}
	initState.Messages = []string{`{
		"str": "foo",
		"num": 100,
		"bool": false,
		"null": null,
		"array": ["foo", "bar", "baz"],
		"obj": { "a": 1, "b": 2 }
	}`}

	return initState
}
