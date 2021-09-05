package main

import (
	"flag"
	"fmt"

	"github.com/gozuk16/gosi"
	"github.com/spf13/nitro"
)

var timer *nitro.B

func main() {
	// 時間計測
	timer = nitro.Initialize()
	flag.Parse()
	timer.Step("init")

	sample1()

	//p := gosi.Process(501)
	//fmt.Println(p)
}

func sample1() {
	fmt.Println(string(gosi.Disk().Json()))
	timer.Step("server stats: disk")
	fmt.Println(string(gosi.Mem().Json()))
	timer.Step("server stats: mem")

	gosi.RefreshCpu()
	fmt.Println(string(gosi.Cpu().Json()))
	timer.Step("server stats: cpu")

	fmt.Println(string(gosi.Info().Json()))
	timer.Step("server stats: info")

	p, err := gosi.CrawlDirs("service", "service.json")
	if err != nil {
		fmt.Println(err)
	}
	timer.Step("CrawlDirs")

	for _, v := range p {
		fmt.Println(v)
	}
	timer.Step("get dirs")
}
