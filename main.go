package main

import (
	"fmt"
	"time"
	"runtime"
	"github.com/leaanthony/mewn"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/wailsapp/wails"
)

type cpustats struct {
	Usage   int
	Count   int
	Os      string
	Arch    string
	CpuInfo []cpu.InfoStat
}

type diskstat struct {
	info *disk.UsageStat
	}
func getStats() cpustats {
	var c cpustats
	stats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	c.CpuInfo = stats
	return c
}

func cpuPercent() []float64{
	p, err := cpu.Percent(time.Second, false)
	if err != nil{
		fmt.Print(err)
	}
	return p
}

func getDiskUsuage() *disk.UsageStat{
	u, err := disk.Usage("/dev/sda5/")
	if err != nil{
		fmt.Println(err)
	}
	return u
}
func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "activity-monitor",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	u, _ := disk.Usage("/")
	fmt.Println(*u)
	info := getStats()
	p := cpuPercent()
	fmt.Println(info)
	fmt.Println(p)
	cp := &cpustats{}
	cp.Os = runtime.GOOS
	cp.Arch = runtime.GOARCH
	fmt.Printf("%v", cp)
//	app.Bind(basic)
	app.Run()
}
