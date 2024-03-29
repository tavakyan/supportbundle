package main

import (
	"fmt"
	"github.com/tavakyan/supportbundle/supportbundle"
	"os"
)

func main() {
	var err error
	// Don't execute if we don't have correct number of arguments
	//if len(os.Args) != 2 {
	//	fmt.Printf("Failed to execute: Invalid number of arguments %v\n", len(os.Args))
	//	fmt.Print("The input should be the path to the bundle archive")
	//	return
	//}
	//
	//We assume the argument in index 1 is the path to the bundle archive
	filePath := os.Args[1]
	//
	//extractedPath, err := supportbundle.ExtractBundle(filePath)
	//if err != nil {
	//	fmt.Printf("Failed to extract %v:", err.Error())
	//	return
	//}

	extractedPath := filePath

	// unmarshal the path into support bundle struct
	supportBundle := supportbundle.SupportBundle{}

	err = supportBundle.UnmarshalBundle(extractedPath)
	if err != nil {
		fmt.Printf("Failed to unmarshal %v:", err.Error())
		return
	}

	supportBundle.PrintInsights()
}
