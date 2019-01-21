package models

import (
	"errors"
	"time"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OS struct {
	ID             int `gorm:"primary_key"`
	UpTime         int64
	StartTime      int64
	UsePermillage  int    //usage, 50 meas 5%
	Arch           string `gorm:"type:varchar(32);not null;"`
	Os             string `gorm:"type:varchar(32);not null;"`
	KernelVersion  string `gorm:"type:varchar(64);not null;"`
	KernelHostname string `gorm:"type:varchar(64);not null;"`
	NumCpu         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SYS struct {
	ID             int    `gorm:"primary_key"`
	Cpu_idle       uint64 // time spent in the idle task
	Cpu_total      uint64 // total of all time fields
	Cpu_permillage int    //usage, 50 meas 5%
	Avg1min        float32
	Avg5min        float32
	Avg15min       float32

	MemTotal          uint64
	MemAvailable      uint64
	MemUsePermillage  int //usage, 50 meas 5%
	SwapTotal         uint64
	SwapFree          uint64
	SwapUsePermillage int //usage, 50 meas 5%

	CreatedAt time.Time
	UpdatedAt time.Time
}

type NETIF struct {
	ID               int `gorm:"primary_key"`
	Iface            string
	InBytes          uint64
	InPackages       uint64
	TotalInBytes     uint64
	TotalInPackages  uint64
	OutBytes         uint64
	OutPackages      uint64
	TotalOutBytes    uint64
	TotalOutPackages uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type NETWORK struct {
	ID             int `gorm:"primary_key"`
	Established    uint64
	TcpConnections uint64
	TcpListen      uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type PARTITION struct {
	ID int `gorm:"primary_key"`

	Device               string
	FsSpec               string
	FsFile               string
	FsVfstype            string
	BytesAll             uint64
	BytesUsed            uint64
	BytesUsedPermillage  int
	InodesAll            uint64
	InodesUsed           uint64
	InodesUsedPermillage int
	ReadRequests         uint64 // Total number of reads completed successfully.
	ReadBytes            uint64 // Total number of Bytes read successfully.
	WriteRequests        uint64 // total number of writes completed successfully.
	WriteBytes           uint64 // total number of Bytes written successfully.

	CreatedAt time.Time
	UpdatedAt time.Time
}

type THERMAL struct {
	ID        int `gorm:"primary_key"`
	Type      string
	Temp      int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *OS) UpdateOS() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func (this *SYS) UpdateSys() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func (this *NETIF) UpdateNETIF() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func (this *NETWORK) UpdateNETWORK() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func (this *PARTITION) UpdatePARTITION() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func (this *THERMAL) UpdateTHERMAL() error {
	this.ID = 0
	db.Save(this)
	return nil
}

func GetOSinfo(limit int) ([]OS, error) {
	var osinfo []OS
	db.Order("id desc").Limit(limit).Find(&osinfo)
	if len(osinfo) < 1 {
		return osinfo, errors.New("not found")
	}
	return osinfo, nil
}

func GetSysInfo(limit int) ([]SYS, error) {
	var sysinfo []SYS
	db.Order("id desc").Limit(limit).Find(&sysinfo)
	if len(sysinfo) < 1 {
		return sysinfo, errors.New("not found")
	}
	return sysinfo, nil
}

func GetNetIfUsage(limit int) ([]NETIF, error) {
	var netifinfo []NETIF
	db.Order("id desc").Limit(limit).Find(&netifinfo)
	if len(netifinfo) < 1 {
		return netifinfo, errors.New("not found")
	}
	return netifinfo, nil
}

func GetNetWorkUsage(limit int) ([]NETWORK, error) {
	var networkinfo []NETWORK
	db.Order("id desc").Limit(limit).Find(&networkinfo)
	if len(networkinfo) < 1 {
		return networkinfo, errors.New("not found")
	}
	return networkinfo, nil
}

func GetDiskUsage(limit int) ([]PARTITION, error) {
	var diskinfo []PARTITION
	db.Order("id desc").Limit(limit).Find(&diskinfo)
	if len(diskinfo) < 1 {
		return diskinfo, errors.New("not found")
	}
	return diskinfo, nil
}

func GetThermal(limit int) ([]THERMAL, error) {
	var thermal []THERMAL
	db.Order("id desc").Limit(limit).Find(&thermal)
	if len(thermal) < 1 {
		return thermal, errors.New("not found")
	}
	return thermal, nil
}
