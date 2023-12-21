package util

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var (
	cpuLogicCount int = 1
)

func init() {
	cpuLogicCount, _ = cpu.Counts(true)
}

type CpuLoadInfo struct {
	CpuNum  int       `json:"cpu_num"`
	LoadAvg []float64 `json:"load_avg"`
}

type MemInfo struct {
	MemorySum     uint64  `json:"memory_sum"`
	MemoryUseRate float64 `json:"memory_use_rate"`
}

func GetCpuLoadInfo() *CpuLoadInfo {
	avgStat, _ := load.Avg()
	info := &CpuLoadInfo{
		CpuNum:  cpuLogicCount,
		LoadAvg: []float64{avgStat.Load1, avgStat.Load5, avgStat.Load15},
	}
	return info
}

func GetMemInfo() *MemInfo {
	v, _ := mem.VirtualMemory()
	info := &MemInfo{
		MemorySum:     v.Total,
		MemoryUseRate: v.UsedPercent,
	}
	return info
}
