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
}

func InitAppState() *AppState {
	var initState = &AppState{}
	initState.ConnectionList = []Bookmark{}
	initState.Messages = []string{}
	initState.AppOptions = DefaultOpts()

	return initState
}

func (a *AppState) SelectServer() {}
