package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	//"gopkg.in/yaml.v2"
)

var keyPrefix = flag.String(
	"key-prefix", "{{.AppName}}", "Template for key prefix on datacenter-specific keys",
)
var globalKeyPrefix = flag.String(
	"global-key-prefix", "", "Template for key prefix on datacenter-specific keys",
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Printf("Usage: %s [options] <input-dir> <output-dir>\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputDir := args[0]
	outputDir := args[1]

	err := run(inputDir, outputDir, *keyPrefix, *globalKeyPrefix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}

func run(inputDir, outputDir, keyPrefix, globalKeyPrefix string) error {

	appDirs, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("error reading input dir: %s", err)
	}

	fmt.Printf("App dirs: %#v\n", appDirs)

	return nil
}
