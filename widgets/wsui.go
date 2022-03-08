package widgets

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

type WSUI struct {
	messageList, connectionList *widget.List
	appState                    *model.AppState
	filter                      *widget.Entry
	connectButton               *widget.Button
	connectionDisplay           []model.Bookmark
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

	ui.messageList = widget.NewList(
		func() int {
			return len(ui.appState.Messages)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i int, o fyne.CanvasObject) {

			o.(*widget.Label).SetText(ui.appState.Messages[i])
		})

	ui.filter = widget.NewEntry()
	ui.filter.OnChanged = ui.SearchConnections
	ui.connectButton = widget.NewButton("Connect", func() {
		log.Println("tapped")
	})

	return container.NewGridWithColumns(3,
		container.NewBorder(container.NewVBox(widget.NewLabel("Connection Bookmarks"), ui.filter), ui.connectButton, nil, nil, container.NewScroll(ui.connectionList)),
		container.NewBorder(container.NewVBox(widget.NewLabel("Messages")), nil, nil, nil, container.NewScroll(ui.messageList)),
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
