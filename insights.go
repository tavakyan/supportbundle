package supportbundle

import (
	"fmt"
	"time"
)

func (sb SupportBundle) getOS() string {
	return "in proc/version"
}

func (sb SupportBundle) getOSVersion() string {
	return "in proc/version"
}

func (sb SupportBundle) getNumCores() string {
	return "in proc/cpuinfo get cpu cores field"
}

func (sb SupportBundle) getLoadAverage(d time.Duration) string {
	return "3rd field from commands/loadavg/loadavg"
}

func (sb SupportBundle) getDiskUsage() string {
	return "default/proc/meminfo MemTotal - MemAvailable"
}

func (sb SupportBundle) getDockerVersion() string {
	return "get from docker_version.json"
}

func (sb SupportBundle) getDockerStorageDriver() string {
	return "gete from docker_info.json"
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
