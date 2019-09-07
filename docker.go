package supportbundle

type DockerInfo struct {
	Driver string
}
type DockerVersion struct {
	Version string
}

type Docker struct {
	DockerVersion DockerVersion
	DockerInfo    DockerInfo
}

func (docker Docker) Unmarshal(path string) error {
	return nil
}
