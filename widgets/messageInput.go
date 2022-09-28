package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MessageInput struct {
	widget.Entry
	lastInput string
}

func NewMessageInput() *MessageInput {
	entry := &MessageInput{}
	entry.ExtendBaseWidget(entry)
	entry.Wrapping = fyne.TextTruncate
	entry.MultiLine = true
	return entry
}

func (e *MessageInput) Clear() {
	e.lastInput = e.Text
	e.SetText("")

}
func (e *MessageInput) TypedKey(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyUp {
		e.SetText(e.lastInput)
	}
	if key.Name == fyne.KeyBackspace {
		if len(e.Text) > 0 {
			e.SetText(e.Text[:len(e.Text)-1])
		}

	}
}
