package main

import (
	"flag"
	"fmt"

	"github.com/gozuk16/goss"
	"github.com/gozuk16/goss/file"
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
	fmt.Println(string(goss.Disk()))
	timer.Step("server stats: disk")
	fmt.Println(string(goss.Mem()))
	timer.Step("server stats: mem")

	goss.RefreshCpu()
	fmt.Println(string(goss.Cpu()))
	timer.Step("server stats: cpu")

	fmt.Println(string(goss.Load()))
	timer.Step("server stats: load")
	fmt.Println(string(goss.Info().Json()))
	timer.Step("server stats: info")

	//d, err := file.CrawlDirs("/Users/gozu/go/src/github.com/gozuk16", "go.mod")
	d, err := file.CrawlDirs("service", "service.json")
	if err != nil {
		fmt.Println(err)
	}
	timer.Step("CrawlDirs")

	for _, v := range d {
		fmt.Println(v)
	}
	timer.Step("get dirs")
}
