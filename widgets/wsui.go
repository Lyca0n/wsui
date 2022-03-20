package widgets

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"

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
	optionsForm       *widget.Form
}

func getOpsFormItems(options *model.Options) []*widget.FormItem {
	var formItems = []*widget.FormItem{}
	val := reflect.Indirect(reflect.ValueOf(options))
	for i := 0; i < val.Type().NumField(); i++ {
		// skips fields without json tag
		if tag, ok := val.Type().Field(i).Tag.Lookup("json"); ok {
			fmt.Println("Afield " + tag)
			formItems = append(formItems, &widget.FormItem{
				Text: tag, Widget: widget.NewEntry()})
		}
	}
	return formItems
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
		ui.connectButton.Enable()
	}
	ui.messageContainer = container.NewVBox(widget.NewLabel("hello this is  a new message"))
	ui.messageScoll = container.NewScroll(ui.messageContainer)
	ui.messageEntry = widget.NewEntry()

	ui.filter = widget.NewEntry()
	ui.filter.OnChanged = ui.SearchConnections
	ui.connectButton = widget.NewButton("Connect", func() {
		ui.handleConnect()
	})
	ui.connectButton.Disable()
	ui.sendButton = widget.NewButton("Send", func() {
		ui.sendHandler(ui.messageEntry.Text)
	})
	ui.sendButton.Disable()
	ui.optionsForm = &widget.Form{
		Items: getOpsFormItems(ui.appState.AppOptions),
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted")

		},
	}
	ui.optionsForm.SubmitText = "Set Configuration"

	return container.NewGridWithColumns(3,
		container.NewBorder(container.NewVBox(widget.NewLabel("Connection Bookmarks"), ui.filter), ui.connectButton, nil, nil, container.NewScroll(ui.connectionList)),
		container.NewBorder(container.NewVBox(widget.NewLabel("Messages")), nil, nil, nil, container.NewVSplit(ui.messageScoll, container.NewBorder(nil, ui.sendButton, nil, nil, ui.messageEntry))),
		container.NewBorder(container.NewVBox(widget.NewLabel("Options")), nil, nil, nil, ui.optionsForm),
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
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			ui.handleReconnect()
		}
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
		log.Print("Error connecting to Websocket Server:", err)
	} else {
		ui.appState.Connection = conn
		go ui.receiveHandler(ui.appState.Connection)

		ui.connectButton.Text = "Disconnect"
		ui.connectButton.OnTapped = func() {
			ui.handleDisconnect()
		}
		ui.connectButton.Refresh()
		ui.sendButton.Enable()
		ui.optionsForm.Disable()
	}
}

func (ui *WSUI) handleDisconnect() {

	ui.appState.Connection.Close()
	ui.connectButton.Text = "Connect"
	ui.connectButton.OnTapped = func() {
		ui.handleConnect()
	}
	ui.connectButton.Refresh()
	ui.sendButton.Disable()
	ui.optionsForm.Enable()
}

func (ui *WSUI) handleReconnect() {

	ui.appState.Connection.Close()
	ui.connectButton.Text = "Reconnect"
	ui.connectButton.OnTapped = func() {
		ui.handleConnect()
	}
	ui.connectButton.Refresh()
}
