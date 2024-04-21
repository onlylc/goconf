package apis

import (
	"fmt"
	"goconf/core/sdk/api"
	"runtime"
	"strconv"
	"time"

	"goconf/core/sdk/pkg"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	netInSpeed, netOutSpeed, netInTransfer, netOutTransfer, lastUpdateNetStats uint64
	cachedBootTime                                                             time.Time
)


type ServerMonitor struct {
	api.Api
}

// GetHourDiffer 获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

func (e ServerMonitor) ServerInfo(c *gin.Context) {
	e.Context = c

	sysInfo, err := host.Info()
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["men"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()
	osDic["ip"] = pkg.GetLocaHonst()
	osDic["projectDir"] = pkg.GetCurrentPath()
	osDic["hostName"] = sysInfo.Hostname
	osDic["time"] = time.Now().Format("2006-01-02 15:04:05")

	memory, _ := mem.VirtualMemory()
	memDic := make(map[string]interface{}, 0)
	memDic["used"] = memory.Used / MB
	memDic["total"] = memory.Total / MB
	memDic["percent"] = pkg.Round(memory.UsedPercent, 2)

	swapDic := make(map[string]interface{}, 0)
	swapDic["used"] = memory.SwapTotal - memory.SwapFree
	swapDic["total"] = memory.SwapTotal

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuInfo"], _ = cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuDic["percent"] = pkg.Round(percent[0], 2)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	disklist := make([]disk.UsageStat, 0)
	//所有分区
	var diskTotal, diskUsed, diskUsedPercent float64
	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				disklist = append(disklist, *diskDetail)

			}
		}
	}

	d, _ := disk.Usage("/")

	diskTotal = float64(d.Total / GB)
	diskUsed = float64(d.Used / GB)
	diskUsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", d.UsedPercent), 64)

	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotal
	diskDic["used"] = diskUsed
	diskDic["percent"] = diskUsedPercent

	bootTime, _ := host.BootTime()
	cachedBootTime = time.Unix(int64(bootTime), 0)

	netDic := make(map[string]interface{}, 0)
	netDic["in"] = pkg.Round(float64(netInSpeed/KB), 2)
	netDic["out"] = pkg.Round(float64(netOutSpeed/KB), 2)
	e.Custom(gin.H{
		"code":     200,
		"os":       osDic,
		"mem":      memDic,
		"cpu":      cpuDic,
		"disk":     diskDic,
		"net":      netDic,
		"swap":     swapDic,
		"location": "Aliyun",
		"bootTime": GetHourDiffer(cachedBootTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")),
	})
}
