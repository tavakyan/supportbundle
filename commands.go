package supportbundle

type LoadAvg struct {
	LoadTime15m string
}

type Commands struct {
	LoadAvg LoadAvg
}

func (commands Commands) Unmarshal(path string) error {

	return nil

}
