package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

type HeaderWidget struct {
	enabled *widget.Check
	name    *widget.Entry
	value   *widget.Entry
}

func (hw *HeaderWidget) Init(option *model.HeaderOption) *fyne.Container {
	hw.enabled = widget.NewCheck("Enabled", func(value bool) {})
	hw.name = widget.NewEntry()
	hw.value = widget.NewEntry()
	hw.name.OnChanged = func(txt string) {
		option.Name = txt
	}
	hw.value.OnChanged = func(txt string) {
		option.Value = txt
	}
	hw.enabled.OnChanged = func(val bool) {
		option.Enabled = val
	}

	hw.name.SetText(option.Name)
	hw.value.SetText(option.Value)
	return container.New(layout.NewGridLayout(3), hw.enabled, hw.name, hw.value)
}

type HeaderListWidget struct {
	list *widget.List
}

func (hlw *HeaderListWidget) Init(headerOptions *[]model.HeaderOption) *widget.List {
	return widget.NewList(
		func() int {
			return len(*headerOptions)
		},
		func() fyne.CanvasObject {
			widget := HeaderWidget{}
			return widget.Init(&(*headerOptions)[len(*headerOptions)-1])
		}, func(lii widget.ListItemID, co fyne.CanvasObject) {},
	)
}
