package system

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/fstanis/screenresolution"
)

type LinuxInfo struct{}

// OS returns the running program's operating system target.
func (LinuxInfo) OS() string { return runtime.GOOS }

// HostName returns the host name reported by the kernel.
func (LinuxInfo) HostName() string {
	host, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return host
}

// Kernel returns the kernel version.
func (LinuxInfo) Kernel() string {
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(out))
}

// Uptime returns the system uptime in seconds.
func (LinuxInfo) Uptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return err.Error()
	}
	parts := strings.Fields(string(data))
	if len(parts) > 0 {
		uptime, err := time.ParseDuration(parts[0] + "s")
		if err != nil {
			return err.Error()
		}
		return uptime.String()
	}
	return "unknown"
}

// Shell returns the user's default shell.
func (LinuxInfo) Shell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	return shell
}

// Resolution returns the primary screen resolution.
func (LinuxInfo) Resolution() string {
	resolution := screenresolution.GetPrimary()
	if resolution == nil {
		return "unknown"
	}
	return resolution.String()
}

// DE returns the desktop environment.
func (LinuxInfo) DE() string {
	de := os.Getenv("XDG_CURRENT_DESKTOP")
	if de == "" {
		return "unknown"
	}
	return de
}

// CPU returns the CPU model name.
func (LinuxInfo) CPU() string {
	cpuinfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		return err.Error()
	}
	// Usually, the first CPU entry has the model name.
	if len(cpuinfo.Processors) > 0 {
		return cpuinfo.Processors[0].ModelName
	}
	return "unknown"
}

// GPU returns the GPU model name.
func (LinuxInfo) GPU() string {
	// This is a simple approach using lspci. A more robust solution might require
	// parsing different outputs or using a dedicated library.
	out, err := exec.Command("lspci").Output()
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA compatible controller") {
			parts := strings.Split(line, ": ")
			if len(parts) > 1 {
				return parts[1]
			}
		}
	}
	return "unknown"
}

// Memory returns the total physical memory.
func (LinuxInfo) Memory() string {
	meminfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return err.Error()
	}
	// MemTotal is in KB, so we convert it to GB.
	totalGB := float64(meminfo.MemTotal) / (1024 * 1024)
	return fmt.Sprintf("%.2f GB", totalGB)
}

