package supportbundle

type OS struct {
	Name    string
	Version string
}

type MemInfo struct {
	MemTotal     string
	MemAvailable string
}

type CpuInfo struct {
	CpuCores string
}

type Proc struct {
	MemInfo MemInfo
	CpuInfo CpuInfo
	OS      OS
}

func (proc Proc) Unmarshal(path string) error {

	return nil

}
