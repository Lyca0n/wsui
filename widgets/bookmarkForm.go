package widgets

import (
	"net/url"

	"fyne.io/fyne/v2/widget"
	"github.com/Lyca0n/wsui/model"
)

var (
	SCHEME_DEFAULT = "ws"
)

type BookmarkForm struct {
	name     *widget.Entry
	hostname *widget.Entry
	scheme   *widget.Entry
	path     *widget.Entry
	form     *widget.Form
}

func (bf *BookmarkForm) Reset() {
	bf.name.Text = ""
	bf.scheme.Text = SCHEME_DEFAULT
	bf.hostname.Text = ""
	bf.path.Text = ""
}

func (bf *BookmarkForm) Init(ui *WSUI) *widget.Form {

	bf.name = widget.NewEntry()
	bf.scheme = widget.NewEntry()
	bf.scheme.Text = SCHEME_DEFAULT
	bf.hostname = widget.NewEntry()
	bf.path = widget.NewEntry()
	bf.form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: bf.name},
			{Text: "Scheme", Widget: bf.scheme},
			{Text: "Hostname", Widget: bf.hostname},
			{Text: "Path", Widget: bf.path},
		},
		OnSubmit: func() {
			ui.AppendBookmark(model.Bookmark{
				Url:  url.URL{Scheme: bf.scheme.Text, Host: bf.hostname.Text, Path: bf.path.Text},
				Name: bf.name.Text,
			})
			bf.Reset()
		},
	}

	return bf.form
}
