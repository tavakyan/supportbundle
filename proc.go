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
