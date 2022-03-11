package widgets

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

type WSUI struct {
	connectionList    *widget.List
	messageScoll      *container.Scroll
	messageContainer  *fyne.Container
	appState          *model.AppState
	filter            *widget.Entry
	connectButton     *widget.Button
	connectionDisplay []model.Bookmark
	messageEntry      *widget.Entry
	sendButton        *widget.Button
}

func (ui *WSUI) MakeUI() fyne.CanvasObject {
	ui.appState = model.InitAppState()
	ui.connectionDisplay = ui.appState.ConnectionList
	ui.connectionList = widget.NewList(
		func() int {
			return len(ui.connectionDisplay)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ui.connectionDisplay[i].String())
		})
	ui.connectionList.OnSelected = func(id widget.ListItemID) {
		fmt.Println("Selected one")
		ui.appState.SelectedServer = &ui.connectionDisplay[id]
	}
	ui.messageContainer = container.NewVBox(widget.NewLabel("hello this is  a new message"))
	ui.messageScoll = container.NewScroll(ui.messageContainer)
	ui.messageEntry = widget.NewEntry()

	ui.filter = widget.NewEntry()
	ui.filter.OnChanged = ui.SearchConnections
	ui.connectButton = widget.NewButton("Connect", func() {
		ui.handleConnect()
	})

	ui.sendButton = widget.NewButton("Send", func() {
		ui.sendHandler(ui.messageEntry.Text)
	})
	return container.NewGridWithColumns(3,
		container.NewBorder(container.NewVBox(widget.NewLabel("Connection Bookmarks"), ui.filter), ui.connectButton, nil, nil, container.NewScroll(ui.connectionList)),
		container.NewBorder(container.NewVBox(widget.NewLabel("Messages")), nil, nil, nil, container.NewVSplit(ui.messageScoll, container.NewBorder(nil, ui.sendButton, nil, nil, ui.messageEntry))),
	)

}

func (ui *WSUI) SearchConnections(term string) {
	fmt.Print("changed " + term + "\n")
	if term != "" {
		ui.connectionDisplay = model.FilterByTerm(term, ui.appState.ConnectionList)
		fmt.Printf("length %d \n", len(ui.appState.ConnectionList))
	} else {
		ui.connectionDisplay = ui.appState.ConnectionList
	}
	ui.connectionList.Refresh()

}

func (ui *WSUI) receiveHandler(connection *websocket.Conn) {
	defer close(done)
	defer ui.appState.Connection.Close()
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)

		ui.appState.Messages = append(ui.appState.Messages, string(msg))
		who := ui.appState.Connection.RemoteAddr()
		messageLabel := widget.NewLabel(who.String() + " : \n" + string(msg))
		ui.appendMessage(messageLabel)
	}
}

func (ui *WSUI) sendHandler(message string) {

	err := ui.appState.Connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error in receive:", err)
		return
	}
	ui.appState.Messages = append(ui.appState.Messages, string(message))
	messageLabel := widget.NewLabel("You : \n" + string(message))
	ui.appendMessage(messageLabel)
}

func (ui *WSUI) appendMessage(newMessage *widget.Label) {
	ui.messageContainer.Objects = append(ui.messageContainer.Objects, newMessage)
	ui.messageContainer.Refresh()
	ui.messageScoll.ScrollToBottom()
}

func (ui *WSUI) handleConnect() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT
	conn, _, err := websocket.DefaultDialer.Dial(ui.appState.SelectedServer.Url.String(), nil)

	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}

	ui.appState.Connection = conn
	go ui.receiveHandler(ui.appState.Connection)

}
