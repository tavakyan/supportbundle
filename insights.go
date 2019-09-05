package supportbundle

import (
	"fmt"
	"time"
)

// Print insights from bundle
func (sb SupportBundle) Print() {
	os := sb.GetOS()
	osVer := sb.GetOSVersion()
	numCores := sb.GetNumCores()
	loadAvg := sb.GetLoadAverage(time.Minute * 15)
	diskUsage := sb.GetDiskUsage()
	dockerVer := sb.GetDockerVersion()
	dockerStorageDrv := sb.GetDockerStorageDriver()

	fmt.Print(os)
	fmt.Print(osVer)
	fmt.Print(numCores)
	fmt.Print(loadAvg)
	fmt.Print(diskUsage)
	fmt.Print(dockerVer)
	fmt.Print(dockerStorageDrv)
}
