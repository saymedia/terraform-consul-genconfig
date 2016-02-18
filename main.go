package main

import (
	"flag"
	"fmt"
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

	fmt.Printf("Input dir is %s\n", inputDir)
	fmt.Printf("Output dir is %s\n", outputDir)
	fmt.Printf("Key prefix is %s\n", *keyPrefix)
	fmt.Printf("Global key prefix is %s\n", *globalKeyPrefix)
}
