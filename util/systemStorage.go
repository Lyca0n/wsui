package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Lyca0n/wsui/model"
	log "github.com/sirupsen/logrus"
)

var (
	LINUX_BOOKMARK_FILE = "wsui.json"
)

func GetUserFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("obtaining user home directory")
	}
	return filepath.Join(home, LINUX_BOOKMARK_FILE)
}

func LoadBookmarks() []model.Bookmark {
	books := []model.Bookmark{}
	content, err := ioutil.ReadFile(GetUserFilePath())
	if err != nil {
		log.Debug(err)
		log.Info("no configuration file found, attempting to create one ")
	}

	json.Unmarshal(content, &books)
	return books
}

func UnloadBookmarks(books []model.Bookmark) {
	empty, err := os.Create(GetUserFilePath())
	if err != nil {
		log.Debug("creating file")
	} else {
		empty.Close()
	}

	literalBooks, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		log.Fatal("writing bookmarks")
	}
	log.Debug(literalBooks)
	if err := os.WriteFile(GetUserFilePath(), literalBooks, 0644); err != nil {
		log.Fatal("writing file ", err)
	}

}
