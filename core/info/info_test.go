package info

import (
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetReportCaller(true)
	os.Exit(m.Run())
}

func TestInit(t *testing.T) {

}


