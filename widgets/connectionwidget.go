package widgets

import (
	"fmt"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

type ConnectionWidget struct {
	listItems   []model.Bookmark
	View        *fyne.Container
	connList    *widget.List
	create      *widget.Entry
	displayList []model.Bookmark
	selected    model.Bookmark
}

func (c *ConnectionWidget) Search(term string) {
	fmt.Print("changed " + term + "\n")
	if term != "" {
		c.displayList = filterByTerm(term, c.listItems)
		fmt.Printf("length %d \n", len(c.displayList))
	} else {
		c.displayList = c.listItems
	}
	c.connList.Refresh()

}

func (c *ConnectionWidget) onSelectConn(id int) {
	c.selected = c.displayList[id]
	fmt.Print(c.selected.Name)
}

func (c *ConnectionWidget) MakeConnectionWidget(connectionList []model.Bookmark) *fyne.Container {
	c.listItems = connectionList
	c.displayList = connectionList
	c.connList = widget.NewList(
		func() int {
			return len(c.displayList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(c.displayList[i].String())
		})

	c.connList.OnSelected = c.onSelectConn
	title := widget.NewLabel("Connection Bookmarks")
	input := widget.NewEntry()

	connectButton := widget.NewButton("Connect", func() {
		log.Println("tapped")
	})
	input.SetPlaceHolder("Search Connections")
	input.OnChanged = c.Search
	cont := container.NewBorder(container.NewVBox(title, input), connectButton, nil, nil,
		container.NewScroll(c.connList))
	c.View = cont
	return c.View
}

func filterByTerm(term string, list []model.Bookmark) []model.Bookmark {
	var data []model.Bookmark
	for _, item := range list {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(term)) {
			data = append(data, item)
		}
	}
	return data
}
