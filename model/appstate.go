package model

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type AppState struct {
	ConnectionList []Bookmark
	Messages       []string
	Connection     *websocket.Conn
	SelectedServer *Bookmark
	AppOptions     *Options
}

func InitAppState() *AppState {
	var initState = &AppState{}
	initState.ConnectionList = []Bookmark{
		{Name: "Home", Url: url.URL{Scheme: "ws", Host: "localhost:9090", Path: "/socket"}},
		{Name: "Store 0020", Url: url.URL{Scheme: "ws", Host: "192.168.0.120", Path: ""}},
	}
	initState.Messages = []string{}
	initState.AppOptions = DefaultOpts()

	return initState
}

func (a *AppState) SelectServer() {}
