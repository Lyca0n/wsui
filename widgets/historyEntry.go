package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type HistoryEntry struct {
	widget.Entry
	lastInput string
}

func NewHistoryEntry() *HistoryEntry {
	entry := &HistoryEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *HistoryEntry) Clear() {
	e.lastInput = e.Text
	e.SetText("")

}
func (e *HistoryEntry) TypedKey(key *fyne.KeyEvent) {
	if key.Name == fyne.KeyUp {
		e.SetText(e.lastInput)
	}
	if key.Name == fyne.KeyBackspace {
		e.SetText(e.Text[:len(e.Text)-1])
	}
}
