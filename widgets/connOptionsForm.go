package widgets

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

var (
	MEDIA_OPTS = []string{"Text", "JSON"}
)

type ConnOptionsForm struct {
	sendAs       *widget.Select
	consumeAs    *widget.Select
	originHeader *widget.Entry
	form         *widget.Form
}

func (cof *ConnOptionsForm) Init(callback func(model.Options)) *widget.Form {
	cof.sendAs = widget.NewSelect(MEDIA_OPTS, func(s string) {})
	cof.consumeAs = widget.NewSelect(MEDIA_OPTS, func(s string) {})
	cof.originHeader = widget.NewEntry()
	cof.form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Send As", Widget: cof.sendAs},
			{Text: "Consume As", Widget: cof.consumeAs},
			{Text: "Origin Header", Widget: cof.originHeader},
		},
		OnSubmit: func() {
			callback(model.Options{
				SendAs:       cof.sendAs.Selected,
				ConsumeAs:    cof.consumeAs.Selected,
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
