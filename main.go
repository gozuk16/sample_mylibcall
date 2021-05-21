package main

import (
	"fmt"

	"github.com/gozuk16/mylib"
)

func main() {
	fmt.Println(string(mylib.Disk()))
	fmt.Println(string(mylib.Mem()))
	fmt.Println(string(mylib.Cpu()))
	fmt.Println(string(mylib.Load()))
}
