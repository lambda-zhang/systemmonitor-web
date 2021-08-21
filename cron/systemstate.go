package monitor

import (
	m "github.com/lambda-zhang/systemmonitor-web/models"

	"os"
	"time"

	s "github.com/lambda-zhang/systemmonitor"
)

var (
	period_sec int = 10
	SM         *s.SysInfo
	Info       AllInfo
)

type AllInfo struct {
	CurOsInfo  oSinfo      `json:"os"`
	Curthermal thermalinfo `json:"thermal"`
	CpuMem     cpumeminfo  `json:"cpumem"`
	Network    networkinfo `json:"network"`
	Netif      netif       `json:"netif"`
	Disk       disk        `json:"disk"`
}

type NetifInfo struct {
	Name              string `json:"name"`
	TotalInKBytes     uint64 `json:"inkb"`
	TotalOutKBytes    uint64 `json:"outkb"`
	TotalInKPackages  uint64 `json:"inkp"`
	TotalOutKPackages uint64 `json:"oukp"`
}

type DiskInfo struct {
	Name       string  `json:"name"`
	MBytesAll  uint64  `json:"totalmb"`
	BytesUsage float32 `json:"usage"`
}

type oSinfo struct {
	UpTime         int64   `json:"uptime"`
	StartTime      int64   `json:"starttime"`
	CpuUsage       float32 `json:"cpuusage"`
	Arch           string  `json:"arch"`
	Os             string  `json:"os"`
	KernelVersion  string  `json:"kernel"`
	KernelHostname string  `json:"hostname"`
	NumCpu         int     `json:"numcpu"`

	MemTotalMB  uint64  `json:"memtotalmb"`
	MemUsage    float32 `json:"memusage"`
	SwapTotalMB uint64  `json:"swaptotalmb"`
	SwapUsage   float32 `json:"swapusage"`

	//netif
	Netinfo []NetifInfo `json:"netinfo"`

	//disk
	DisksInfo []DiskInfo `json:"diskinfo"`

	UpdatedAt int64 `json:"-"`
}

func (_osinfo *oSinfo) Updateinfo(info *s.SysInfo) error {
	o := info.OS
	_osinfo.UpTime = o.UpTime
	_osinfo.StartTime = o.StartTime * 1000
	_osinfo.CpuUsage = float32(o.UsePermillage) / 10
	_osinfo.Arch = o.Arch
	_osinfo.Os = o.Os
	_osinfo.KernelVersion = o.KernelVersion
	_osinfo.KernelHostname = o.KernelHostname
	_osinfo.NumCpu = o.NumCPU

	mem := info.Mem
	_osinfo.MemTotalMB = mem.MemTotal / 1024 / 1024
	_osinfo.MemUsage = float32(mem.MemUsePermillage) / 10
	_osinfo.SwapTotalMB = mem.SwapTotal / 1024 / 1024
	_osinfo.SwapUsage = float32(mem.SwapUsePermillage) / 10

	_osinfo.Netinfo = make([]NetifInfo, 0)
	n := info.Net.Cards
	for _, v := range n {
		_osinfo.Netinfo = append(_osinfo.Netinfo, NetifInfo{
			Name:              v.Iface,
			TotalInKBytes:     v.TotalInBytes / 1024,
			TotalOutKBytes:    v.TotalOutBytes / 1024,
			TotalInKPackages:  v.TotalInPackages / 1024,
			TotalOutKPackages: v.TotalOutPackages / 1024,
		})
	}
	_osinfo.DisksInfo = make([]DiskInfo, 0)
	disk := info.Fs.Disks
	for _, v := range disk {
		_osinfo.DisksInfo = append(_osinfo.DisksInfo, DiskInfo{
			Name:       v.Device,
			MBytesAll:  v.BytesAll / 1024 / 1024,
			BytesUsage: float32(v.BytesUsedPermillage) / 10,
		})
	}
	_osinfo.UpdatedAt = time.Now().Unix() * 1000
	return nil
}

type _thermalinfo struct {
	Name string  `json:"name"`
	Temp float32 `json:"val"`
}

type thermalinfo struct {
	UpdatedAt int64          `json:"updateat"`
	Temp      []_thermalinfo `json:"temp"`
}

func (_thermal *thermalinfo) Updateinfo(info *s.SysInfo) error {
	t := info.Thermal
	_thermal.Temp = make([]_thermalinfo, 0)
	_thermal.UpdatedAt = time.Now().Unix() * 1000
	for _, v := range t.Thermal {
		_thermal.Temp = append(_thermal.Temp, _thermalinfo{Name: v.Type, Temp: float32(v.Temp) / 1000})
	}
	return nil
}

type cpumeminfo struct {
	CpuUsage  float32 `json:"cpuusage"`
	Avg1min   float32 `json:"avg1min"`
	Avg5min   float32 `json:"avg5min"`
	Avg15min  float32 `json:"avg15min"`
	MemUsage  float32 `json:"memusage"`
	SwapUsage float32 `json:"swapusage"`
	UpdatedAt int64   `json:"updateat"`
}

func (_cpumem *cpumeminfo) Updateinfo(info *s.SysInfo) error {
	c := info.CPU
	mem := info.Mem

	_cpumem.UpdatedAt = time.Now().Unix() * 1000
	_cpumem.CpuUsage = float32(c.CPUs["cpu"].CPUPermillage) / 10
	_cpumem.Avg1min = c.Avg1min
	_cpumem.Avg5min = c.Avg5min
	_cpumem.Avg15min = c.Avg15min
	_cpumem.MemUsage = float32(mem.MemUsePermillage) / 10
	_cpumem.SwapUsage = float32(mem.SwapUsePermillage) / 10
	return nil
}

type networkinfo struct {
	Established    uint64 `json:"tcpest"`
	TcpConnections uint64 `json:"tcpall"`
	TcpListen      uint64 `json:"tcplis"`
	UpdatedAt      int64  `json:"updateat"`
}

func (_network *networkinfo) Updateinfo(info *s.SysInfo) error {
	n := info.Net

	_network.UpdatedAt = time.Now().Unix() * 1000
	_network.Established = n.TCP.TCPEstablished + n.TCP6.TCPEstablished
	_network.TcpConnections = n.TCP.TCPConnections + n.TCP6.TCPConnections
	_network.TcpListen = n.TCP.TCPListen + n.TCP6.TCPListen
	return nil
}

type _netif struct {
	Name        string `json:"name"`
	InKBytes    uint64 `json:"inkb"`
	InPackages  uint64 `json:"inp"`
	OutKBytes   uint64 `json:"outkb"`
	OutPackages uint64 `json:"outp"`
}

type netif struct {
	UpdatedAt int64    `json:"updateat"`
	NetIf     []_netif `json:"cards"`
}

func (__netif *netif) Updateinfo(info *s.SysInfo) error {
	n := info.Net
	__netif.NetIf = make([]_netif, 0)

	for _, v := range n.Cards {
		__netif.NetIf = append(__netif.NetIf, _netif{
			Name:        v.Iface,
			InKBytes:    v.InBytes / uint64(period_sec) / 1024,
			OutKBytes:   v.OutBytes / uint64(period_sec) / 1024,
			InPackages:  v.InPackages / uint64(period_sec),
			OutPackages: v.OutPackages / uint64(period_sec),
		})
	}
	__netif.UpdatedAt = time.Now().Unix() * 1000
	return nil
}

type _disk struct {
	Name      string  `json:"name"`
	ReadKBps  uint64  `json:"rkbps"`
	WriteKBps uint64  `json:"wkbps"`
	ReadRps   uint64  `json:"rrps"`
	WriteRps  uint64  `json:"wrps"`
	Usage     float32 `json:"usage"`
}

type disk struct {
	UpdatedAt int64   `json:"updateat"`
	Disks     []_disk `json:"disks"`
}

func (__disk *disk) Updateinfo(info *s.SysInfo) error {
	f := info.Fs
	__disk.Disks = make([]_disk, 0)

	for _, v := range f.Disks {
		__disk.Disks = append(__disk.Disks, _disk{
			Name:      v.Device,
			ReadKBps:  v.ReadBytes / 1024 / uint64(period_sec),
			WriteKBps: v.WriteBytes / 1024 / uint64(period_sec),
			ReadRps:   v.ReadRequests / uint64(period_sec),
			WriteRps:  v.WriteRequests / uint64(period_sec),
			Usage:     float32(v.BytesUsedPermillage) / 10,
		})
	}

	__disk.UpdatedAt = time.Now().Unix() * 1000
	return nil
}

func callback(sysinfo *s.SysInfo) {
	o := sysinfo.OS
	c := sysinfo.CPU
	mem := sysinfo.Mem
	n := sysinfo.Net
	f := sysinfo.Fs
	t := sysinfo.Thermal
	ts := time.Now()

	Info.CurOsInfo.Updateinfo(sysinfo)
	Info.Curthermal.Updateinfo(sysinfo)
	Info.CpuMem.Updateinfo(sysinfo)
	Info.Network.Updateinfo(sysinfo)
	Info.Netif.Updateinfo(sysinfo)
	Info.Disk.Updateinfo(sysinfo)

	usedb := os.Getenv("USEDB")
	if usedb != "true" {
		return
	}

	osdb := &m.OS{ID: 0, UpTime: o.UpTime, StartTime: o.StartTime, UsePermillage: o.UsePermillage,
		Arch: o.Arch, Os: o.Os, KernelVersion: o.KernelVersion,
		KernelHostname: o.KernelHostname, NumCpu: o.NumCPU, UpdatedAt: ts}
	osdb.UpdateOS()

	sysdb := &m.SYS{Cpu_idle: c.CPUs["cpu"].CPUIdle, Cpu_total: c.CPUs["cpu"].CPUTotal, Cpu_permillage: c.CPUs["cpu"].CPUPermillage,
		Avg1min: c.Avg1min, Avg5min: c.Avg5min, Avg15min: c.Avg15min, MemTotal: mem.MemTotal,
		MemAvailable: mem.MemAvailable, MemUsePermillage: mem.MemUsePermillage,
		SwapTotal: mem.SwapTotal, SwapFree: mem.SwapFree, SwapUsePermillage: mem.SwapUsePermillage,
		UpdatedAt: ts}

	sysdb.UpdateSys()

	for _, v := range n.Cards {
		netifdb := &m.NETIF{Iface: v.Iface, InBytes: v.InBytes / uint64(period_sec),
			InPackages: v.InPackages / uint64(period_sec), TotalInBytes: v.TotalInBytes,
			TotalInPackages: v.TotalInPackages, OutBytes: v.OutBytes / uint64(period_sec),
			OutPackages:   v.OutPackages / uint64(period_sec),
			TotalOutBytes: v.TotalOutBytes, TotalOutPackages: v.TotalOutPackages, UpdatedAt: ts}
		netifdb.UpdateNETIF()
	}

	networkdb := &m.NETWORK{TcpConnections: n.TCP.TCPClosing + n.TCP6.TCPConnections,
		Established: n.TCP.TCPEstablished + n.TCP6.TCPEstablished,
		TcpListen:   n.TCP.TCPListen + n.TCP6.TCPListen, UpdatedAt: ts}
	networkdb.UpdateNETWORK()

	for _, v := range f.Disks {
		partitiondb := &m.PARTITION{Device: v.Device, FsSpec: v.FsSpec, FsFile: v.FsFile,
			FsVfstype: v.FsVfstype, BytesAll: v.BytesAll, BytesUsed: v.BytesUsed,
			BytesUsedPermillage: v.BytesUsedPermillage, InodesAll: v.InodesAll,
			InodesUsed: v.InodesUsed, InodesUsedPermillage: v.InodesUsedPermillage,
			ReadRequests: v.ReadRequests, ReadBytes: v.ReadBytes, WriteRequests: v.WriteRequests,
			WriteBytes: v.WriteBytes, UpdatedAt: ts}
		partitiondb.UpdatePARTITION()
	}

	for _, v := range t.Thermal {
		thermaldb := &m.THERMAL{Type: v.Type, Temp: v.Temp}
		thermaldb.UpdateTHERMAL()
	}
}

func init() {
	SM = s.New(period_sec, callback)
	SM.OSEn = true
	SM.CPUEn = true
	SM.MemEn = true
	SM.NetEn = true
	SM.FsEn = true
	SM.ThermalEn = true
	SM.Start()

	// defer sm.Stop()
}
