package info

import (
  "os"
  "runtime"
  log "github.com/sirupsen/logrus"
  "github.com/shirou/gopsutil/v3/disk"
)

type AppInfo struct {
  os string
  arch string
  hostname string
  homeDir string
  disks []string
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
  Info.os = runtime.GOOS
  Info.arch = runtime.GOARCH
  hostname, err := os.Hostname()
  homeDir, err := os.UserHomeDir()
  Info.hostname = hostname
  Info.homeDir = homeDir
  return err
}

func (appInfo *AppInfo) Print() {
  log.Infof("Gathered information:")
  log.Infof("OS       : %s", Info.os);
  log.Infof("ARCH     : %s", Info.arch);
  log.Infof("Hostname : %s", Info.hostname);
  log.Infof("Home dir : %s", Info.homeDir);
}
