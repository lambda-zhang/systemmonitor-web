package controllers

import (
	m "github.com/lambda-zhang/systemmonitor-web/models"

	"time"

	"github.com/gin-gonic/gin"
)

func checkUniq(a []int64, val int64) int {
	for _, v := range a {
		if val == v {
			return 1
		}
	}
	return 0
}

func Max(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func MaxInt(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

type OSinfo struct {
	UpTime         int64
	StartTime      int64
	UsePermillage  int
	Arch           string
	Os             string
	KernelVersion  string
	KernelHostname string
	NumCpu         int

	MemTotalMB  uint64
	MemUsage    float32
	SwapTotalMB uint64
	SwapUsage   float32

	//netif
	TotalInKBytes    map[string]uint64
	TotalOutKBytes   map[string]uint64
	TotalInPackages  map[string]uint64
	TotalOutPackages map[string]uint64

	//disk
	MBytesAll  map[string]uint64
	BytesUsage map[string]float32

	UpdatedAt time.Time
}

func _GetOsInfo() OSinfo {
	var info OSinfo
	osinfo, err := m.GetOSinfo(1)
	mem, err2 := m.GetSysInfo(1)
	netif, err3 := m.GetNetIfUsage(10)
	disk, err4 := m.GetDiskUsage(10)
	if err == nil && err2 == nil && err3 == nil && err4 == nil {
		info.TotalInKBytes = make(map[string]uint64, 0)
		info.TotalOutKBytes = make(map[string]uint64, 0)
		info.TotalInPackages = make(map[string]uint64, 0)
		info.TotalOutPackages = make(map[string]uint64, 0)
		info.MBytesAll = make(map[string]uint64, 0)
		info.BytesUsage = make(map[string]float32, 0)
		for _, v := range netif {
			if _, ok := info.TotalInKBytes[v.Iface]; !ok {
				info.TotalInKBytes[v.Iface] = v.TotalInBytes / 1024
			}
			if _, ok := info.TotalOutKBytes[v.Iface]; !ok {
				info.TotalOutKBytes[v.Iface] = v.TotalOutBytes / 1024
			}
			if _, ok := info.TotalInPackages[v.Iface]; !ok {
				info.TotalInPackages[v.Iface] = v.TotalInPackages
			}
			if _, ok := info.TotalOutPackages[v.Iface]; !ok {
				info.TotalOutPackages[v.Iface] = v.TotalOutPackages
			}
		}
		for _, v := range disk {
			if _, ok := info.MBytesAll[v.Device]; !ok {
				info.MBytesAll[v.Device] = v.BytesAll / 1024 / 1024
			}
			if _, ok := info.BytesUsage[v.Device]; !ok {
				info.BytesUsage[v.Device] = float32(v.BytesUsedPermillage) / 10
			}
		}

		info.UpTime = osinfo[0].UpTime * 1000
		info.StartTime = osinfo[0].StartTime * 1000
		info.UsePermillage = osinfo[0].UsePermillage
		info.Arch = osinfo[0].Arch
		info.Os = osinfo[0].Os
		info.KernelVersion = osinfo[0].KernelVersion
		info.KernelHostname = osinfo[0].KernelHostname
		info.NumCpu = osinfo[0].NumCpu

		info.MemTotalMB = mem[0].MemTotal / 1024 / 1024
		info.MemUsage = float32(mem[0].MemUsePermillage) / 10
		info.SwapTotalMB = mem[0].SwapTotal / 1024 / 1024
		info.SwapUsage = float32(mem[0].SwapUsePermillage) / 10
	}
	return info
}

func GetOsInfo(c *gin.Context) {
	var info OSinfo
	osinfo, err := m.GetOSinfo(1)
	mem, err2 := m.GetSysInfo(1)
	netif, err3 := m.GetNetIfUsage(10)
	disk, err4 := m.GetDiskUsage(10)
	if err == nil && err2 == nil && err3 == nil && err4 == nil {
		info.TotalInKBytes = make(map[string]uint64, 0)
		info.TotalOutKBytes = make(map[string]uint64, 0)
		info.TotalInPackages = make(map[string]uint64, 0)
		info.TotalOutPackages = make(map[string]uint64, 0)
		info.MBytesAll = make(map[string]uint64, 0)
		info.BytesUsage = make(map[string]float32, 0)
		for _, v := range netif {
			if _, ok := info.TotalInKBytes[v.Iface]; !ok {
				info.TotalInKBytes[v.Iface] = v.TotalInBytes / 1024
			}
			if _, ok := info.TotalOutKBytes[v.Iface]; !ok {
				info.TotalOutKBytes[v.Iface] = v.TotalOutBytes / 1024
			}
			if _, ok := info.TotalInPackages[v.Iface]; !ok {
				info.TotalInPackages[v.Iface] = v.TotalInPackages
			}
			if _, ok := info.TotalOutPackages[v.Iface]; !ok {
				info.TotalOutPackages[v.Iface] = v.TotalOutPackages
			}
		}
		for _, v := range disk {
			if _, ok := info.MBytesAll[v.Device]; !ok {
				info.MBytesAll[v.Device] = v.BytesAll / 1024 / 1024
			}
			if _, ok := info.BytesUsage[v.Device]; !ok {
				info.BytesUsage[v.Device] = float32(v.BytesUsedPermillage) / 10
			}
		}

		info.UpTime = osinfo[0].UpTime * 1000
		info.StartTime = osinfo[0].StartTime * 1000
		info.UsePermillage = osinfo[0].UsePermillage
		info.Arch = osinfo[0].Arch
		info.Os = osinfo[0].Os
		info.KernelVersion = osinfo[0].KernelVersion
		info.KernelHostname = osinfo[0].KernelHostname
		info.NumCpu = osinfo[0].NumCpu

		info.MemTotalMB = mem[0].MemTotal / 1024 / 1024
		info.MemUsage = float32(mem[0].MemUsePermillage) / 10
		info.SwapTotalMB = mem[0].SwapTotal / 1024 / 1024
		info.SwapUsage = float32(mem[0].SwapUsePermillage) / 10

		info.UpdatedAt = osinfo[0].UpdatedAt
		c.JSON(200, gin.H{"status": 200, "data": info})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}

type cpumemstate struct {
	UpdatedAt       []int64
	UpdatedAtNewest int64
	Cpu_percent     []float32
	Avg1min         []float32
	Avg5min         []float32
	Avg15min        []float32
	MaxLoad         float32
	MemUsepercent   []float32 //usage, 50 meas 5%
	SwapUsepercent  []float32 //usage, 50 meas 5%
	NumCpu          int
}

func GetCpuState(c *gin.Context) {
	var cmstate cpumemstate
	osinfo, err1 := m.GetOSinfo(1)
	cpu, err := m.GetSysInfo(50)
	if err == nil && err1 == nil {
		cmstate.UpdatedAt = make([]int64, 0)
		cmstate.Cpu_percent = make([]float32, 0)
		cmstate.Avg1min = make([]float32, 0)
		cmstate.Avg5min = make([]float32, 0)
		cmstate.Avg15min = make([]float32, 0)
		cmstate.MemUsepercent = make([]float32, 0)
		cmstate.SwapUsepercent = make([]float32, 0)
		cmstate.UpdatedAtNewest = cpu[0].UpdatedAt.Unix() * 1000
		cmstate.NumCpu = osinfo[0].NumCpu
		for i := len(cpu) - 1; i >= 0; i-- {
			cmstate.UpdatedAt = append(cmstate.UpdatedAt, cpu[i].UpdatedAt.Unix()*1000)
			cmstate.Cpu_percent = append(cmstate.Cpu_percent, float32(cpu[i].Cpu_permillage)/10)
			cmstate.Avg1min = append(cmstate.Avg1min, cpu[i].Avg1min)
			cmstate.Avg5min = append(cmstate.Avg5min, cpu[i].Avg5min)
			cmstate.Avg15min = append(cmstate.Avg15min, cpu[i].Avg15min)
			cmstate.MemUsepercent = append(cmstate.MemUsepercent, float32(cpu[i].MemUsePermillage)/10)
			cmstate.SwapUsepercent = append(cmstate.SwapUsepercent, float32(cpu[i].SwapUsePermillage)/10)
		}
		c.JSON(200, gin.H{"status": 200, "data": cmstate})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}

type netifusage struct {
	UpdatedAt       []int64
	UpdatedAtNewest int64
	InKBytes        map[string][]uint64
	OutKBytes       map[string][]uint64
	MaxKBytes       uint64
	InPackages      map[string][]uint64
	OutPackages     map[string][]uint64
	MaxPackages     uint64
}

func GetNetState(c *gin.Context) {
	var usage netifusage
	netif, err := m.GetNetIfUsage(90)
	if err == nil {
		usage.UpdatedAt = make([]int64, 0)
		usage.InKBytes = make(map[string][]uint64, 0)
		usage.OutKBytes = make(map[string][]uint64, 0)
		usage.InPackages = make(map[string][]uint64, 0)
		usage.OutPackages = make(map[string][]uint64, 0)
		usage.UpdatedAtNewest = netif[0].UpdatedAt.Unix() * 1000
		usage.MaxKBytes = 5
		usage.MaxPackages = 10
		for i := len(netif) - 1; i >= 0; i-- {
			v := netif[i]
			if _, ok := usage.InKBytes[v.Iface]; !ok {
				usage.InKBytes[v.Iface] = make([]uint64, 0)
			}
			if _, ok := usage.OutKBytes[v.Iface]; !ok {
				usage.OutKBytes[v.Iface] = make([]uint64, 0)
			}
			if _, ok := usage.InPackages[v.Iface]; !ok {
				usage.InPackages[v.Iface] = make([]uint64, 0)
			}
			if _, ok := usage.OutPackages[v.Iface]; !ok {
				usage.OutPackages[v.Iface] = make([]uint64, 0)
			}

			usage.MaxKBytes = Max(usage.MaxKBytes, v.InBytes/1024)
			usage.MaxKBytes = Max(usage.MaxKBytes, v.OutBytes/1024)
			usage.MaxPackages = Max(usage.MaxPackages, v.InPackages)
			usage.MaxPackages = Max(usage.MaxPackages, v.OutPackages)

			usage.InKBytes[v.Iface] = append(usage.InKBytes[v.Iface], v.InBytes/1024)
			usage.OutKBytes[v.Iface] = append(usage.OutKBytes[v.Iface], v.OutBytes/1024)
			usage.InPackages[v.Iface] = append(usage.InPackages[v.Iface], v.InPackages)
			usage.OutPackages[v.Iface] = append(usage.OutPackages[v.Iface], v.OutPackages)

			isUniq := checkUniq(usage.UpdatedAt, v.UpdatedAt.Unix()*1000)
			if isUniq == 0 {
				usage.UpdatedAt = append(usage.UpdatedAt, v.UpdatedAt.Unix()*1000)
			}
		}
		c.JSON(200, gin.H{"status": 200, "data": usage})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}

type tcpusage struct {
	UpdatedAt         []int64
	UpdatedAtNewest   int64
	TcpConnections    []uint64
	Established       []uint64
	TcpListen         []uint64
	MaxTcpConnections uint64
}

func GetTcpState(c *gin.Context) {
	var usage tcpusage
	tcp, err := m.GetNetWorkUsage(50)
	if err == nil {
		usage.UpdatedAt = make([]int64, 0)
		usage.TcpConnections = make([]uint64, 0)
		usage.Established = make([]uint64, 0)
		usage.TcpListen = make([]uint64, 0)
		usage.UpdatedAtNewest = tcp[0].UpdatedAt.Unix() * 1000
		for i := len(tcp) - 1; i >= 0; i-- {
			usage.UpdatedAt = append(usage.UpdatedAt, tcp[i].UpdatedAt.Unix()*1000)
			usage.TcpConnections = append(usage.TcpConnections, tcp[i].TcpConnections)
			usage.Established = append(usage.Established, tcp[i].Established)
			usage.MaxTcpConnections = Max(usage.MaxTcpConnections, tcp[i].TcpConnections)
			usage.TcpListen = append(usage.TcpListen, tcp[i].TcpListen)
		}
		c.JSON(200, gin.H{"status": 200, "data": usage})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}

type diskstate struct {
	UpdatedAt            []int64
	UpdatedAtNewest      int64
	ReadKBytes           map[string][]uint64
	WriteKBytes          map[string][]uint64
	MaxKBytes            uint64
	ReadRequests         map[string][]uint64
	WriteRequests        map[string][]uint64
	MaxRequests          uint64
	InodesUsedPermillage map[string][]float32
	BytesUsedPermillage  map[string][]float32
}

func GetDiskState(c *gin.Context) {
	var state diskstate
	disk, err := m.GetDiskUsage(90)
	if err == nil {
		state.UpdatedAt = make([]int64, 0)
		state.ReadKBytes = make(map[string][]uint64, 0)
		state.WriteKBytes = make(map[string][]uint64, 0)
		state.MaxKBytes = 0
		state.ReadRequests = make(map[string][]uint64, 0)
		state.WriteRequests = make(map[string][]uint64, 0)
		state.MaxRequests = 0
		state.InodesUsedPermillage = make(map[string][]float32, 0)
		state.BytesUsedPermillage = make(map[string][]float32, 0)
		state.UpdatedAtNewest = disk[0].UpdatedAt.Unix() * 1000
		for i := len(disk) - 1; i >= 0; i-- {
			v := disk[i]
			if _, ok := state.ReadKBytes[v.Device]; !ok {
				state.ReadKBytes[v.Device] = make([]uint64, 0)
			}
			if _, ok := state.WriteKBytes[v.Device]; !ok {
				state.WriteKBytes[v.Device] = make([]uint64, 0)
			}
			if _, ok := state.ReadRequests[v.Device]; !ok {
				state.ReadRequests[v.Device] = make([]uint64, 0)
			}
			if _, ok := state.WriteRequests[v.Device]; !ok {
				state.WriteRequests[v.Device] = make([]uint64, 0)
			}
			if _, ok := state.InodesUsedPermillage[v.Device]; !ok {
				state.InodesUsedPermillage[v.Device] = make([]float32, 0)
			}
			if _, ok := state.BytesUsedPermillage[v.Device]; !ok {
				state.BytesUsedPermillage[v.Device] = make([]float32, 0)
			}

			state.MaxRequests = Max(state.MaxRequests, v.ReadRequests)
			state.MaxRequests = Max(state.MaxRequests, v.WriteRequests)
			state.MaxKBytes = Max(state.MaxKBytes, v.ReadBytes/1024)
			state.MaxKBytes = Max(state.MaxKBytes, v.WriteBytes/1024)

			isUniq := checkUniq(state.UpdatedAt, v.UpdatedAt.Unix()*1000)
			if isUniq == 0 {
				state.UpdatedAt = append(state.UpdatedAt, v.UpdatedAt.Unix()*1000)
			}
			state.ReadKBytes[v.Device] = append(state.ReadKBytes[v.Device], v.ReadBytes/1024)
			state.WriteKBytes[v.Device] = append(state.WriteKBytes[v.Device], v.WriteBytes/1024)
			state.ReadRequests[v.Device] = append(state.ReadRequests[v.Device], v.ReadRequests)
			state.WriteRequests[v.Device] = append(state.WriteRequests[v.Device], v.WriteRequests)
			state.InodesUsedPermillage[v.Device] = append(state.InodesUsedPermillage[v.Device], float32(v.InodesUsedPermillage)/10)
			state.BytesUsedPermillage[v.Device] = append(state.BytesUsedPermillage[v.Device], float32(v.BytesUsedPermillage)/10)
		}
		c.JSON(200, gin.H{"status": 200, "data": state})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}

type thermalstate struct {
	UpdatedAt       []int64
	UpdatedAtNewest int64
	Temp            map[string][]float32
	MaxTemp         float32
}

func GetThermalState(c *gin.Context) {
	var state thermalstate
	var maxtemp int64
	state.UpdatedAt = make([]int64, 0)
	state.Temp = make(map[string][]float32, 0)
	state.MaxTemp = 1
	temp, err := m.GetThermal(50)
	if err == nil {
		state.UpdatedAtNewest = temp[0].UpdatedAt.Unix() * 1000
		for i := len(temp) - 1; i >= 0; i-- {
			v := temp[i]
			if _, ok := state.Temp[v.Type]; !ok {
				state.Temp[v.Type] = make([]float32, 0)
			}
			maxtemp = MaxInt(maxtemp, v.Temp)
			isUniq := checkUniq(state.UpdatedAt, v.UpdatedAt.Unix()*1000)
			if isUniq == 0 {
				state.UpdatedAt = append(state.UpdatedAt, v.UpdatedAt.Unix()*1000)
			}
			state.Temp[v.Type] = append(state.Temp[v.Type], float32(v.Temp-v.Temp%100)/1000)
		}
		state.MaxTemp = float32(maxtemp-maxtemp%1000)/1000 + 1
		c.JSON(200, gin.H{"status": 200, "data": state})
	} else {
		c.JSON(200, gin.H{"status": 200, "data": ""})
	}
}
