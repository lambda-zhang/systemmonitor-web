package monitor

import (
	m "../models"

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

func (this *oSinfo) Updateinfo(info *s.SysInfo) error {
	o := info.OS
	this.UpTime = o.UpTime
	this.StartTime = o.StartTime * 1000
	this.CpuUsage = float32(o.UsePermillage) / 10
	this.Arch = o.Arch
	this.Os = o.Os
	this.KernelVersion = o.KernelVersion
	this.KernelHostname = o.KernelHostname
	this.NumCpu = o.NumCpu

	mem := info.Mem
	this.MemTotalMB = mem.MemTotal / 1024 / 1024
	this.MemUsage = float32(mem.MemUsePermillage) / 10
	this.SwapTotalMB = mem.SwapTotal / 1024 / 1024
	this.SwapUsage = float32(mem.SwapUsePermillage) / 10

	this.Netinfo = make([]NetifInfo, 0)
	n := info.Net.Cards
	for _, v := range n {
		this.Netinfo = append(this.Netinfo, NetifInfo{
			Name:              v.Iface,
			TotalInKBytes:     v.TotalInBytes / 1024,
			TotalOutKBytes:    v.TotalOutBytes / 1024,
			TotalInKPackages:  v.TotalInPackages / 1024,
			TotalOutKPackages: v.TotalOutPackages / 1024,
		})
	}
	this.DisksInfo = make([]DiskInfo, 0)
	disk := info.Fs.Disks
	for _, v := range disk {
		this.DisksInfo = append(this.DisksInfo, DiskInfo{
			Name:       v.Device,
			MBytesAll:  v.BytesAll / 1024 / 1024,
			BytesUsage: float32(v.BytesUsedPermillage) / 10,
		})
	}
	this.UpdatedAt = time.Now().Unix() * 1000
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

func (this *thermalinfo) Updateinfo(info *s.SysInfo) error {
	t := info.Thermal
	this.Temp = make([]_thermalinfo, 0)
	this.UpdatedAt = time.Now().Unix() * 1000
	for _, v := range t.Thermal {
		this.Temp = append(this.Temp, _thermalinfo{Name: v.Type, Temp: float32(v.Temp) / 1000})
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

func (this *cpumeminfo) Updateinfo(info *s.SysInfo) error {
	c := info.CPU
	mem := info.Mem

	this.UpdatedAt = time.Now().Unix() * 1000
	this.CpuUsage = float32(c.Cpu_permillage) / 10
	this.Avg1min = c.Avg1min
	this.Avg5min = c.Avg5min
	this.Avg15min = c.Avg15min
	this.MemUsage = float32(mem.MemUsePermillage) / 10
	this.SwapUsage = float32(mem.SwapUsePermillage) / 10
	return nil
}

type networkinfo struct {
	Established    uint64 `json:"tcpest"`
	TcpConnections uint64 `json:"tcpall"`
	TcpListen      uint64 `json:"tcplis"`
	UpdatedAt      int64  `json:"updateat"`
}

func (this *networkinfo) Updateinfo(info *s.SysInfo) error {
	n := info.Net

	this.UpdatedAt = time.Now().Unix() * 1000
	this.Established = n.Tcp.Tcp_established + n.Tcp6.Tcp_established
	this.TcpConnections = n.Tcp.TcpConnections + n.Tcp6.TcpConnections
	this.TcpListen = n.Tcp.Tcp_listen + n.Tcp6.Tcp_listen
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

func (this *netif) Updateinfo(info *s.SysInfo) error {
	n := info.Net
	this.NetIf = make([]_netif, 0)

	for _, v := range n.Cards {
		this.NetIf = append(this.NetIf, _netif{
			Name:        v.Iface,
			InKBytes:    v.InBytes / uint64(period_sec) / 1024,
			OutKBytes:   v.OutBytes / uint64(period_sec) / 1024,
			InPackages:  v.InPackages / uint64(period_sec),
			OutPackages: v.OutPackages / uint64(period_sec),
		})
	}
	this.UpdatedAt = time.Now().Unix() * 1000
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

func (this *disk) Updateinfo(info *s.SysInfo) error {
	f := info.Fs
	this.Disks = make([]_disk, 0)

	for _, v := range f.Disks {
		this.Disks = append(this.Disks, _disk{
			Name:      v.Device,
			ReadKBps:  v.ReadBytes / 1024 / uint64(period_sec),
			WriteKBps: v.WriteBytes / 1024 / uint64(period_sec),
			ReadRps:   v.ReadRequests / uint64(period_sec),
			WriteRps:  v.WriteRequests / uint64(period_sec),
			Usage:     float32(v.BytesUsedPermillage) / 10,
		})
	}

	this.UpdatedAt = time.Now().Unix() * 1000
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
		KernelHostname: o.KernelHostname, NumCpu: o.NumCpu, UpdatedAt: ts}
	osdb.UpdateOS()

	sysdb := &m.SYS{Cpu_idle: c.Cpu_idle, Cpu_total: c.Cpu_total, Cpu_permillage: c.Cpu_permillage,
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

	networkdb := &m.NETWORK{TcpConnections: n.Tcp.TcpConnections + n.Tcp6.TcpConnections,
		Established: n.Tcp.Tcp_established + n.Tcp6.Tcp_established,
		TcpListen:   n.Tcp.Tcp_listen + n.Tcp6.Tcp_listen, UpdatedAt: ts}
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
	SM.Start()
}
