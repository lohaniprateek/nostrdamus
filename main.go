package main

import (
	"fmt"
	"runtime"

	"github.com/lohaniprateek/nostradamus/system"
)

type Info interface {
	OS() string
	HostName() string
	Kernel() string
	Uptime() string
	Shell() string
	Resolution() string
	DE() string
	CPU() string
	GPU() string
	Memory() string
}

func main() {
	var sys Info
	switch runtime.GOOS {
	case "linux":
		sys = system.LinuxInfo{}
	case "windows":
	// sys = system.WindowsInfo{}
	case "darwin":
	//	sys = system.DarwinInfo{}
	default:
		fmt.Println("Unsupported OS: ", runtime.GOOS)
		return
	}

	fmt.Println("OS:", sys.OS())
	fmt.Println("HostName:", sys.HostName())
	fmt.Println("Kernel:", sys.Kernel())
	fmt.Println("Uptime:", sys.Uptime())
	fmt.Println("Shell:", sys.Shell())
	fmt.Println("Resolution:", sys.Resolution())
	fmt.Println("DE:", sys.DE())
	fmt.Println("CPU:", sys.CPU())
	fmt.Println("GPU:", sys.GPU())
	fmt.Println("Memory:", sys.Memory())
}
