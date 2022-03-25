package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Alert struct {
	alertModal   *widget.PopUp
	alertLabel   *widget.Label
	alertConfirm *widget.Button
}

func (a *Alert) makeAlert(win *fyne.Window) {
	a.alertLabel = widget.NewLabel("")
	a.alertConfirm = widget.NewButton("Ok", func() {
		a.onConfirm()
	})
	a.alertModal = widget.NewModalPopUp(container.NewHBox(a.alertLabel, a.alertConfirm), (*win).Canvas())
}

func (a *Alert) onConfirm() {
	a.alertLabel.Text = ""
	a.alertLabel.Refresh()
	a.alertModal.Hide()
}

func (a *Alert) Alert(alert string) {
	a.alertLabel.Text = alert
	a.alertLabel.Refresh()
	a.alertModal.Show()
}
