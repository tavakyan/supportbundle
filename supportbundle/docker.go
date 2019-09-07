package supportbundle

import (
	"encoding/json"
	"io/ioutil"
)

const DockerVersionPath = "/docker_version.json"
const DockerInfoPath = "/docker_info.json"
const DockerPath = "/docker"

type DockerVersion struct {
	Version string
}

func (dv *DockerVersion) unmarshal(path string) error {

	// Set path to LoadAvg command
	// an load file from that path
	p := path + DockerVersionPath
	res, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, dv)
	if err != nil {
		return err
	}

	return nil
}

type DockerInfo struct {
	Driver string
}

func (di *DockerInfo) unmarshal(path string) error {
	// Set path to LoadAvg command
	// an load file from that path
	p := path + DockerInfoPath
	res, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, di)
	if err != nil {
		return err
	}

	return nil
}

type Docker struct {
	DockerVersion DockerVersion
	DockerInfo    DockerInfo
}

func (docker *Docker) Unmarshal(path string) (err error) {

	err = docker.DockerVersion.unmarshal(path + DockerPath)
	if err != nil {
		return err
	}

	err = docker.DockerInfo.unmarshal(path + DockerPath)
	if err != nil {
		return err
	}

	return nil
}
