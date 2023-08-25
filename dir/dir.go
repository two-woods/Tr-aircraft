package dir

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func ListALlDir(path string) []os.DirEntry {
	var err error
	var dirList []os.DirEntry
	dirList, err = os.ReadDir(path)
	if err != nil {
		log.Errorf("%+v", err)
	}
	return dirList
}
