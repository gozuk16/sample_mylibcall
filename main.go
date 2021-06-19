package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gozuk16/goss"
	"github.com/gozuk16/goss/file"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/nitro"
)

var timer *nitro.B

func main() {
	// 時間計測
	timer = nitro.Initialize()
	flag.Parse()
	timer.Step("init")

	// Ctrl+C(割り込み)を受け取る
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	done := make(chan error, 1)
	go sampleServ(done)

	done2 := make(chan error, 1)
	go cpuPread(done2)

	select {
	case <-quit:
		fmt.Println("Interrup signal accepted.")
	case err := <-done:
		fmt.Println("exit.", err)
	}
}

func cpuPread(done2 chan<- error) {
	for {
		fmt.Println(string(goss.Cpu()))
		time.Sleep(1 * time.Second)
	}
	done2 <- nil
	close(done2)
}
func sampleServ(done chan<- error) {
	for {
		goss.RefreshCpu()
		time.Sleep(5 * time.Second)
	}
	done <- nil
	close(done)
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

// getPid 渡されたサービス名のPIDファイルからPIDを取得して返す。エラーなら-1を返す
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
