package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Lyca0n/wsui/widgets"
)

func main() {
	u := &widgets.WSUI{}
	a := app.New()
	w := a.NewWindow("WSUI")
	w.Resize(fyne.NewSize(960, 660))
	w.SetContent(u.MakeUI())
	w.ShowAndRun()
}
