package fs

import (
	"testing"
	"strings"
	"path/filepath"
	"os"
	log "github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
}

func TestWalk(t *testing.T) {
	root, _ := os.UserHomeDir()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(path)) == ".pdf" {
			log.Infof("Found PDF file : %v in %v.", filepath.Dir(path), info.Name())
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}

}
