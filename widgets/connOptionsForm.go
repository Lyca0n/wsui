package widgets

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

type ConnOptionsForm struct {
	sendAs       *widget.Select
	consumeAs    *widget.Select
	originHeader *widget.Entry
	form         *widget.Form
}

func (cof *ConnOptionsForm) Init(callback func(model.Options)) *widget.Form {
	var opts = []string{model.JSON.String(), model.Text.String()}
	cof.sendAs = widget.NewSelect(opts, func(s string) {})
	cof.consumeAs = widget.NewSelect(opts, func(s string) {})
	cof.originHeader = widget.NewEntry()
	cof.form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Send As", Widget: cof.sendAs},
			{Text: "Consume As", Widget: cof.consumeAs},
			{Text: "Origin Header", Widget: cof.originHeader},
		},
		OnSubmit: func() {
			callback(model.Options{
				SendAs:       model.ValueFromString(cof.sendAs.Selected),
				ConsumeAs:    model.ValueFromString(cof.consumeAs.Selected),
				OriginHeader: cof.originHeader.Text,
			})
		},
	}
	return cof.form
}

func (cof *ConnOptionsForm) Disable() {
	cof.form.Disable()
}

func (cof *ConnOptionsForm) Enable() {
	cof.form.Enable()
}
