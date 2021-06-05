package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gozuk16/goss"
	"github.com/gozuk16/goss/file"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/nitro"
)

func main() {
	timer := nitro.Initialize()
	flag.Parse()
	timer.Step("init")

	fmt.Println(string(goss.Disk()))
	timer.Step("server stats: disk")
	fmt.Println(string(goss.Mem()))
	timer.Step("server stats: mem")
	fmt.Println(string(goss.Cpu()))
	timer.Step("server stats: cpu")
	fmt.Println(string(goss.Load()))
	timer.Step("server stats: load")
	fmt.Println(string(goss.Info()))
	timer.Step("server stats: info")

	//d, err := file.CrawlDirs("/Users/gozu/go/src/github.com/gozuk16", "go.mod")
	d, err := file.CrawlDirs("service", "service.json")
	if err != nil {
		fmt.Println(err)
	}
	timer.Step("CrawlDirs")

	var wg sync.WaitGroup
	for _, v := range d {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			getPid(v)
		}(v)
	}
	wg.Wait()
	timer.Step("get pid with goroutine")

	for _, v := range d {
		fmt.Println(string(goss.Process(getPid(v))))
	}
	timer.Step("get process")

	var wg2 sync.WaitGroup
	for _, v := range d {
		wg2.Add(1)
		go func(v string) {
			defer wg2.Done()
			fmt.Println(string(goss.Process(getPid(v))))
		}(v)
	}
	wg2.Wait()
	timer.Step("get process with goroutine")
}

func getPid(service string) int32 {
	f, err := os.Open(filepath.Join("/var/run", service+".pid"))
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return -1
	}
	scanner := bufio.NewScanner(f)
	var t string
	for scanner.Scan() {
		t = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return -1
	}

	pidno, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return -1
	}
	isAlive, err := process.PidExists(int32(pidno))
	fmt.Printf("%v: %d, %v\n", service, pidno, isAlive)
	return int32(pidno)
}
