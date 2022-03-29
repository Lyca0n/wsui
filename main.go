package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Lyca0n/wsui/util"
	"github.com/Lyca0n/wsui/widgets"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(util.GetLogLevel())
	storedBookmarks := util.LoadBookmarks()
	u := &widgets.WSUI{}
	a := app.New()
	w := a.NewWindow("WSUI")
	w.Resize(fyne.NewSize(960, 660))
	w.SetContent(u.MakeUI(&w, storedBookmarks))
	w.ShowAndRun()
}
