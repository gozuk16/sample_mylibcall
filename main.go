package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/gozuk16/goss"
	"github.com/gozuk16/goss/file"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	fmt.Println(string(goss.Disk()))
	fmt.Println(string(goss.Mem()))
	fmt.Println(string(goss.Cpu()))
	fmt.Println(string(goss.Load()))
	fmt.Println(string(goss.Info()))
	//d, err := file.CrawlDirs("/Users/gozu/go/src/github.com/gozuk16", "go.mod")
	d, err := file.CrawlDirs("service", "service.json")
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	for _, v := range d {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			getPid(v)
		}(v)
	}
	wg.Wait()

}

func getPid(service string) {
	f, err := os.Open(filepath.Join("/var/run", service+".pid"))
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return
	}
	scanner := bufio.NewScanner(f)
	var t string
	for scanner.Scan() {
		t = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return
	}

	pidno, err := strconv.Atoi(t)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%v: %d\n", service, -1)
		return
	}
	isAlive, err := process.PidExists(int32(pidno))
	fmt.Printf("%v: %d, %v\n", service, pidno, isAlive)
	return
}
