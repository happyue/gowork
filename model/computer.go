package model

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// ComputerInfo 计算机详细信息
type ComputerInfo struct {
	InfoStat          *host.InfoStat        `json:"InfoStat"`
	CpuInfo           *CpuInfo               `json:"CpuInfo"`
	VirtualMemoryStat *mem.VirtualMemoryStat `json:"VirtualMemoryStat"`
	UsageStat         *disk.UsageStat         `json:"UsageStat"`
}

// type DiskStatus struct {
// 	Name string  `json:"Name"`
// 	All  float64 `json:"All"`
// 	Used float64 `json:"Used"`
// 	Free float64 `json:"Free"`
// }

// const (
// 	B  = 1
// 	KB = 1024 * B
// 	MB = 1024 * KB
// 	GB = 1024 * MB
// )

// type UsageStat struct {
//     Path              string  `json:"path"`
//     Fstype            string  `json:"fstype"`
//     Total             uint64  `json:"total"`
//     Free              uint64  `json:"free"`
//     Used              uint64  `json:"used"`
//     UsedPercent       float64 `json:"usedPercent"`
//     InodesTotal       uint64  `json:"inodesTotal"`
//     InodesUsed        uint64  `json:"inodesUsed"`
//     InodesFree        uint64  `json:"inodesFree"`
//     InodesUsedPercent float64 `json:"inodesUsedPercent"`
// }

type CpuInfo struct {
	Counts  int     `json:"Counts"`
	Percent []float64 `json:"Percent"`
}

// type InfoStat struct {
//     Hostname             string `json:"hostname"`
//     Uptime               uint64 `json:"uptime"`
//     BootTime             uint64 `json:"bootTime"`
//     Procs                uint64 `json:"procs"`           // number of processes
//     OS                   string `json:"os"`              // ex: freebsd, linux
//     Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
//     PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
//     PlatformVersion      string `json:"platformVersion"` // version of the complete OS
//     KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
//     KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
//     VirtualizationSystem string `json:"virtualizationSystem"`
//     VirtualizationRole   string `json:"virtualizationRole"` // guest or host
//     HostID               string `json:"hostid"`             // ex: uuid
// }

// type VirtualMemoryStat struct {
//     // Total amount of RAM on this system
//     Total uint64 `json:"total"`

//     // RAM available for programs to allocate
//     //
//     // This value is computed from the kernel specific values.
//     Available uint64 `json:"available"`

//     // RAM used by programs
//     //
//     // This value is computed from the kernel specific values.
//     Used uint64 `json:"used"`

//     // Percentage of RAM used by programs
//     //
//     // This value is computed from the kernel specific values.
//     UsedPercent float64 `json:"usedPercent"`

//     // This is the kernel's notion of free memory; RAM chips whose bits nobody
//     // cares about the value of right now. For a human consumable number,
//     // Available is what you really want.
//     Free uint64 `json:"free"`

//     // OS X / BSD specific numbers:
//     // http://www.macyourself.com/2010/02/17/what-is-free-wired-active-and-inactive-system-memory-ram/
//     Active   uint64 `json:"active"`
//     Inactive uint64 `json:"inactive"`
//     Wired    uint64 `json:"wired"`

//     // FreeBSD specific numbers:
//     // https://reviews.freebsd.org/D8467
//     Laundry uint64 `json:"laundry"`

//     // Linux specific numbers
//     // https://www.centos.org/docs/5/html/5.1/Deployment_Guide/s2-proc-meminfo.html
//     // https://www.kernel.org/doc/Documentation/filesystems/proc.txt
//     // https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
//     Buffers        uint64 `json:"buffers"`
//     Cached         uint64 `json:"cached"`
//     Writeback      uint64 `json:"writeback"`
//     Dirty          uint64 `json:"dirty"`
//     WritebackTmp   uint64 `json:"writebacktmp"`
//     Shared         uint64 `json:"shared"`
//     Slab           uint64 `json:"slab"`
//     SReclaimable   uint64 `json:"sreclaimable"`
//     SUnreclaim     uint64 `json:"sunreclaim"`
//     PageTables     uint64 `json:"pagetables"`
//     SwapCached     uint64 `json:"swapcached"`
//     CommitLimit    uint64 `json:"commitlimit"`
//     CommittedAS    uint64 `json:"committedas"`
//     HighTotal      uint64 `json:"hightotal"`
//     HighFree       uint64 `json:"highfree"`
//     LowTotal       uint64 `json:"lowtotal"`
//     LowFree        uint64 `json:"lowfree"`
//     SwapTotal      uint64 `json:"swaptotal"`
//     SwapFree       uint64 `json:"swapfree"`
//     Mapped         uint64 `json:"mapped"`
//     VMallocTotal   uint64 `json:"vmalloctotal"`
//     VMallocUsed    uint64 `json:"vmallocused"`
//     VMallocChunk   uint64 `json:"vmallocchunk"`
//     HugePagesTotal uint64 `json:"hugepagestotal"`
//     HugePagesFree  uint64 `json:"hugepagesfree"`
//     HugePageSize   uint64 `json:"hugepagesize"`
// }
