package supportbundle

import (
	"bufio"
	"os"
	"strings"
)

const OSPath = "/version"
const CPUInfoPath = "/cpuinfo"
const MemInfoPath = "/meminfo"
const ProcPath = "/proc"

type OS struct {
	Name    string
	Version string
}

func (os OS) unmarshal(path string) error {
	return nil
}

type MemInfo struct {
	MemTotal     string
	MemAvailable string
}

func (memInfo MemInfo) unmarshal(path string) error {

	p := path + MemInfoPath
	f, err := os.Open(p)
	defer f.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(f)

	var line string
	var parsed map[string]string

	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		res := strings.Split(line, ":")

		k, v := res[0], res[1]

		parsed[k] = v
	}

	memInfo.MemAvailable = parsed["MemAvailable"]
	memInfo.MemTotal = parsed["MemTotal"]

	return nil
}

type CpuInfo struct {
	CpuCores string
}

func (ci CpuInfo) unmarshal(path string) error {
	return nil
}

type Proc struct {
	MemInfo MemInfo
	CpuInfo CpuInfo
	OS      OS
}

func (proc Proc) Unmarshal(path string) (err error) {

	err = proc.MemInfo.unmarshal(path + ProcPath)
	if err != nil {
		return err
	}

	err = proc.CpuInfo.unmarshal(path + ProcPath)
	if err != nil {
		return err
	}

	err = proc.OS.unmarshal(path + ProcPath)
	if err != nil {
		return err
	}

	return nil

}
