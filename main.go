package main

import (
	"flag"
	"fmt"

	"github.com/gozuk16/goss"
	"github.com/spf13/nitro"
)

var timer *nitro.B

func main() {
	// 時間計測
	timer = nitro.Initialize()
	flag.Parse()
	timer.Step("init")

	sample1()

	//p := goss.Process(501)
	//fmt.Println(p)
}

func sample1() {
	fmt.Println(string(goss.Disk().Json()))
	timer.Step("server stats: disk")
	fmt.Println(string(goss.Mem().Json()))
	timer.Step("server stats: mem")

	goss.RefreshCpu()
	fmt.Println(string(goss.Cpu().Json()))
	timer.Step("server stats: cpu")

	fmt.Println(string(goss.Info().Json()))
	timer.Step("server stats: info")

	p, err := goss.CrawlDirs("service", "service.json")
	if err != nil {
		fmt.Println(err)
	}
	timer.Step("CrawlDirs")

	for _, v := range p {
		fmt.Println(v)
	}
	timer.Step("get dirs")
}
