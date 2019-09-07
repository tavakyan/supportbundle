package main

import (
	"fmt"
	"github.com/tavakyan/supportbundle/supportbundle"
	"os"
	"testing"
)

// TODO:
// Write integration test using supportbundle_ex
//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//if err != nil {
//log.Fatal(err)
//}
//extractedPath := dir + "/supportbundle_ex"

// Hack integration test
func TestMain_Integration(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := pwd + "/supportbundle_ex"

	// unmarshal the path into support bundle struct
	supportBundle := supportbundle.SupportBundle{}

	err = supportBundle.UnmarshalBundle(path)
	if err != nil {
		fmt.Printf("Failed to unmarshal %v:", err.Error())
		return
	}

	if supportBundle.Commands.LoadAvg.LoadTime15m != "0.05" {
		t.Error("Failed to parse load avg")
	}

	if supportBundle.Docker.DockerInfo.Driver != "overlay2" {
		t.Error("Failed to parse docker info")
	}

	if supportBundle.Docker.DockerVersion.Version != "18.09.6" {
		t.Error("failed to parse docker version")
	}

	supportBundle.PrintInsights()
}
