package widgets

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
	"github.com/Lyca0n/wsui/util"
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
	messageEntry      *MessageInput
	sendButton        *widget.Button
	optionsForm       *ConnOptionsForm
	newConnButton     *widget.Button
	delConnButton     *widget.Button
	newConnForm       *BookmarkForm
	newConnModal      *widget.PopUp
	alertPopup        *Alert
}

func (ui *WSUI) MakeUI(win *fyne.Window, storedBookmarks []model.Bookmark) fyne.CanvasObject {
	ui.appState = model.InitAppState()
	ui.appState.ConnectionList = storedBookmarks
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
	ui.newConnButton = widget.NewButton("+", func() {
		ui.newConnModal.Show()
	})

	ui.delConnButton = widget.NewButton("-", func() {
		ui.DeleteActiveBookmark()
	})
	ui.delConnButton.Disable()

	ui.connectionList.OnSelected = func(id widget.ListItemID) {
		ui.appState.SelectedServer = &ui.connectionDisplay[id]
		ui.connectButton.Enable()
		ui.delConnButton.Enable()
	}
	ui.messageContainer = container.NewVBox()
	ui.messageScoll = container.NewScroll(ui.messageContainer)
	ui.messageEntry = NewMessageInput()
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
	ui.optionsForm = &ConnOptionsForm{}

	//bookmark Modal
	ui.newConnForm = &BookmarkForm{}
	ui.newConnModal = widget.NewModalPopUp(ui.newConnForm.Init(ui.AppendBookmark), (*win).Canvas())
	ui.newConnModal.Resize(fyne.NewSize(320, 280))

	ui.alertPopup = &Alert{}
	ui.alertPopup.makeAlert(win)
	return container.NewGridWithColumns(3,
		container.NewBorder(container.NewVBox(container.NewHBox(widget.NewLabel("Connection Bookmarks"), ui.newConnButton, ui.delConnButton), ui.filter), ui.connectButton, nil, nil, container.NewScroll(ui.connectionList)),
		container.NewBorder(container.NewVBox(widget.NewLabel("Messages")), nil, nil, nil, container.NewVSplit(ui.messageScoll, container.NewBorder(nil, ui.sendButton, nil, nil, ui.messageEntry))),
		container.NewBorder(container.NewVBox(widget.NewLabel("Options")), nil, nil, nil, ui.optionsForm.Init(ui.setOptions)),
	)

}

func (ui *WSUI) setOptions(newOpts model.Options) {
	fmt.Print("changed opts")
	ui.appState.AppOptions = &newOpts
}

func (ui *WSUI) AppendBookmark(bookmark model.Bookmark) {
	ui.appState.ConnectionList = append(ui.appState.ConnectionList, bookmark)
	ui.connectionDisplay = ui.appState.ConnectionList
	ui.connectionList.Refresh()
	ui.newConnModal.Hide()
	util.UnloadBookmarks(ui.appState.ConnectionList)
}

func (ui *WSUI) DeleteActiveBookmark() {
	fmt.Print("deleting")
	name := ui.appState.SelectedServer.Name
	var filtered []model.Bookmark
	for i := range ui.appState.ConnectionList {
		if ui.appState.ConnectionList[i].Name != name {
			filtered = append(filtered, ui.appState.ConnectionList[i])
		}
	}
	ui.appState.ConnectionList = filtered
	ui.connectionDisplay = ui.appState.ConnectionList
	ui.connectionList.Refresh()
	util.UnloadBookmarks(ui.appState.ConnectionList)
}

func (ui *WSUI) SearchConnections(term string) {
	if term != "" {
		ui.connectionDisplay = model.FilterByTerm(term, ui.appState.ConnectionList)
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

		msgStr := ui.formatMessage(msg)
		ui.appState.Messages = append(ui.appState.Messages, ui.formatMessage(msg))
		who := ui.appState.Connection.RemoteAddr()
		messageLabel := widget.NewLabel(who.String() + " : \n" + msgStr)
		ui.appendMessage(messageLabel)
	}
}

func (ui *WSUI) formatMessage(msg []byte) string {
	msgStr := ""
	if ui.appState.AppOptions.ConsumeAs == model.JSON {

		msgStr = util.FormatJson(msg)
	} else {
		msgStr = util.FormatDefault(msg)
	}

	return msgStr
}

func (ui *WSUI) sendHandler(message string) {

	err := ui.appState.Connection.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error in receive:", err)
		return
	}

	msgStr := ui.formatMessage([]byte(message))
	ui.appState.Messages = append(ui.appState.Messages, string(msgStr))
	messageLabel := widget.NewLabel("You : \n" + string(msgStr))
	ui.appendMessage(messageLabel)
	ui.messageEntry.Clear()
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
	headers := http.Header{}
	headers.Set("Origin", ui.appState.AppOptions.OriginHeader)
	conn, _, err := websocket.DefaultDialer.Dial(ui.appState.SelectedServer.Url.String(), headers)

	if err != nil {
		ui.alertPopup.Alert(err.Error())
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
