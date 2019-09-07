package supportbundle

import (
	"fmt"
	"strconv"
	"time"
)

func (sb SupportBundle) getOS() string {
	return sb.Proc.OS.Name
}

func (sb SupportBundle) getOSVersion() string {
	return sb.Proc.OS.Version
}

func (sb SupportBundle) getNumCores() string {
	return sb.Proc.CpuInfo.CpuCores
}

// getLoadAverage returns the 15 minute cpu load time
// which is the 3rd field in the Load avg command
func (sb SupportBundle) getLoadAverage(d time.Duration) string {
	return sb.Commands.LoadAvg.LoadTime15m
}

// getDiskUsage returns the difference between
// total & available memory
func (sb SupportBundle) getDiskUsage() string {
	memTotal := sb.Proc.MemInfo.MemTotal
	memAvail := sb.Proc.MemInfo.MemAvailable

	memTotalInt, err := strconv.Atoi(memTotal)
	if err != nil {
		return ""
	}

	memTotalAvailInt, err := strconv.Atoi(memAvail)
	if err != nil {
		return ""
	}

	memUsage := memTotalInt - memTotalAvailInt
	return strconv.Itoa(memUsage)
}

func (sb SupportBundle) getDockerVersion() string {
	return sb.Docker.DockerVersion.Version
}

func (sb SupportBundle) getDockerStorageDriver() string {
	return sb.Docker.DockerInfo.Driver
}

// Print insights from bundle
func (sb SupportBundle) PrintInsights() {
	os := sb.getOS()
	osVer := sb.getOSVersion()
	numCores := sb.getNumCores()
	loadAvg := sb.getLoadAverage(time.Minute * 15)
	diskUsage := sb.getDiskUsage()
	dockerVer := sb.getDockerVersion()
	dockerStorageDrv := sb.getDockerStorageDriver()

	fmt.Println("Printing insights for bundle")
	fmt.Println(os)
	fmt.Println(osVer)
	fmt.Println(numCores)
	fmt.Println(loadAvg)
	fmt.Println(diskUsage)
	fmt.Println(dockerVer)
	fmt.Println(dockerStorageDrv)
}
