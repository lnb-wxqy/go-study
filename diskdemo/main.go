package main

import (
	context2 "context"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"time"
)

type Disk struct {
	Total       uint64  `json:"total"` //磁盘总容量
	Use         uint64  `json:"use"`   //磁盘已使用大小
	Free        uint64  `json:"free"`  // 剩余磁盘大小
	UsedPercent float64 `json:"usedPercent"`
}

// 获取根路径大小
func DiskCounts(path string) *Disk {
	ctx, cancel := context2.WithTimeout(context2.Background(), time.Duration(50)*time.Second)
	defer cancel()
	usageStat, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		fmt.Println("获取磁盘失败", err)
		return &Disk{}
	}
	return &Disk{
		Total:       usageStat.Total,
		Use:         usageStat.Used,
		Free:        usageStat.Free,
		UsedPercent: usageStat.UsedPercent,
	}
}

func main() {
	rootMem := DiskCounts("/")
	fmt.Println("rootMem: ", rootMem)
	homeMem := DiskCounts("/home")
	fmt.Println("homeMem: ", homeMem)

}
