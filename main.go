package main

import (
	"fmt"
	"github.com/leaanthony/mewn"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/wailsapp/wails"
	"runtime"
	"time"
)

type stat struct {
	log *wails.CustomLogger
}

type cpustats struct {
	Usage   float64
	Count   int
	Os      string
	Arch    string
	CpuInfo []cpu.InfoStat
}

type diskstat struct {
	info *disk.UsageStat
}

func (s *stat) getAllStats() cpustats {
	return cpustats{
		Usage:   s.cpuPercent(),
		Count:   s.getCount(),
		Os:      runtime.GOOS,
		Arch:    runtime.GOARCH,
		CpuInfo: s.getCpuInfo(),
	}
}

func (s *stat) getCpuInfo() []cpu.InfoStat {
	stats, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	return stats
}

func (s *stat) cpuPercent() float64 {
	p, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Print(err)
	}
	return p[0]
}

func (s *stat) getCount() int {
	count, err := cpu.Counts(true)
	if err != nil {
		s.log.Errorf("Unable to get cpu usage: %s", err)
	}
	return count
}
func getDiskUsuage() *disk.UsageStat {
	u, err := disk.Usage("/dev/sda5/")
	if err != nil {
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

	//initialize cpu calls
	var (
		info = &stat{}
		c    cpustats
	)

	c = info.getAllStats()
	fmt.Println(c)

	u, _ := disk.Usage("/")
	fmt.Println(*u)
	//	app.Bind(basic)
	app.Run()
}
