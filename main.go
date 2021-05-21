package main

import (
	"fmt"

	"github.com/gozuk16/goss"
)

func main() {
	fmt.Println(string(goss.Disk()))
	fmt.Println(string(goss.Mem()))
	fmt.Println(string(goss.Cpu()))
	fmt.Println(string(goss.Load()))
}
