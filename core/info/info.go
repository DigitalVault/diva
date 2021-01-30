package info

import (
  "os"
  "runtime"
  log "github.com/sirupsen/logrus"
  "github.com/shirou/gopsutil/v3/disk"
)

type AppInfo struct {
  Os string
  Arch string
  Hostname string
  HomeDir string
  Disks []string
}

var Info AppInfo

func init() {
  Info.Gather()
  Info.SysInfo()
  //Info.Print()
}

func (appInfo *AppInfo) SysInfo() (error) {
  partitions, err := disk.Partitions(true)
  for i, p:= range partitions {
    log.Debug(i, p)
  }
  return err
}

func (appInfo *AppInfo) Gather() (error) {
  log.Debug("Gathering information about the computer.")
  Info.Os = runtime.GOOS
  Info.Arch = runtime.GOARCH
  hostname, err := os.Hostname()
  homeDir, err := os.UserHomeDir()
  Info.Hostname = hostname
  Info.HomeDir = homeDir
  return err
}

func (appInfo *AppInfo) Print() {
  log.Infof("Gathered information:")
  log.Infof("OS       : %s", Info.Os);
  log.Infof("ARCH     : %s", Info.Arch);
  log.Infof("Hostname : %s", Info.Hostname);
  log.Infof("Home dir : %s", Info.HomeDir);
}
