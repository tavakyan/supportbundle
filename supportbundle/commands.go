package supportbundle

import (
	"bufio"
	"os"
	"strings"
)

const LoadAvgPath = "/loadavg/loadavg"
const CommandPath = "/commands"

type LoadAvg struct {
	LoadTime15m string
}

// Converts file containing:
// 0.26 0.14 0.05 5/233 5186
// to LoadAvg structure, error otherwise
func (avg *LoadAvg) unmarshal(path string) error {

	// Set path to LoadAvg command
	// an load file from that path
	p := path + LoadAvgPath
	f, err := os.Open(p)
	defer f.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')

	// Delimit by spaces
	res := strings.Split(line, " ")

	// Set LoadAvg params
	avg.LoadTime15m = res[2]

	return nil
}

type Commands struct {
	LoadAvg LoadAvg
}

func (commands Commands) Unmarshal(path string) (err error) {

	err = commands.LoadAvg.unmarshal(path + CommandPath)
	if err != nil {
		return err
	}

	return nil

}
