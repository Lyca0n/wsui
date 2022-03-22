package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Lyca0n/wsui/model"
)

var (
	LINUX_BOOKMARK_FILE = "./bookmarks.json"
	LINUX_BOOKMARK_DIR  = "./"
)

func LoadBookmarks() []model.Bookmark {
	books := []model.Bookmark{}
	content, err := ioutil.ReadFile(LINUX_BOOKMARK_FILE)
	if err != nil {
		fmt.Print("Error when opening file: ", err)
	}

	json.Unmarshal(content, &books)
	return books
}

func UnloadBookmarks(books []model.Bookmark) {
	empty, err := os.Create(LINUX_BOOKMARK_FILE)
	if err != nil {
		fmt.Print("Error creating file")
	} else {
		empty.Close()
	}

	literalBooks, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		fmt.Print("Unrecoverable Error Writing bookmarks")
	}
	fmt.Print(literalBooks)
	if err := os.WriteFile(LINUX_BOOKMARK_FILE, literalBooks, 0644); err != nil {
		fmt.Print("Write Error ", err)
	}

}
