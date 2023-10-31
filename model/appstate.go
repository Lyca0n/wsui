package model

import (
	"github.com/gorilla/websocket"
)

type AppState struct {
	ConnectionList []Bookmark
	Messages       []string
	Connection     *websocket.Conn
	SelectedServer *Bookmark
	AppOptions     *Options
	Headers        []HeaderOption
}

func InitAppState() *AppState {
	var initState = &AppState{}
	initState.ConnectionList = []Bookmark{}
	initState.Messages = []string{}
	initState.AppOptions = DefaultOpts()
	initState.Headers = []HeaderOption{{
		Enabled: false,
		Name:    "Origin",
		Value:   "http://localhost",
	}}
	return initState
}

func (a *AppState) SelectServer() {}
