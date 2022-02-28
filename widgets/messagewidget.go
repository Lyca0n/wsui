package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MessageWidget struct {
	messages    []string
	messageList *widget.List
}

func (mw *MessageWidget) MakeMessageWidget() *fyne.Container {
	str := `{
		"str": "foo",
		"num": 100,
		"bool": false,
		"null": null,
		"array": ["foo", "bar", "baz"],
		"obj": { "a": 1, "b": 2 }
	}`

	mw.messages = []string{str}
	mw.messageList = widget.NewList(
		func() int {
			return len(mw.messages)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(mw.messages[i])
		})
	cont := container.NewBorder(container.NewVBox(widget.NewLabel("Messages")), nil, nil, nil,
		container.NewScroll(mw.messageList))

	return cont
}

func (mw *MessageWidget) NewMessage(message string) {
	mw.messages = append(mw.messages, message)
	mw.messageList.Refresh()
}
