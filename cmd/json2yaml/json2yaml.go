package main

import (
	"io/ioutil"
	"log"
	"os"

	"sigs.k8s.io/yaml"
)

func main() {
	j, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to read from stdin: %v", err)
	}
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		log.Fatalf("unable to convert json into yaml: %v", err)
	}
	_, err = os.Stdout.Write(y)
	if err != nil {
		log.Fatalf("unable to write to stdout: %v", err)
	}
}
