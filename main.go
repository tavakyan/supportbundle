package supportbundle

import (
	"fmt"
)

type SupportBundle struct {
}

func main() {

	extractedPath, err := ExtractBundle(filePath)
	if err != nil {
		return
	}

	// load the file into high level structure
	supportBundle, err := UnmarshalBundle(extractedPath)

	// print the structure
	fmt.Print(supportBundle)
}
