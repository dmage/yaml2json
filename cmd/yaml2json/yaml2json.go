package main

import (
	"io/ioutil"
	"log"
	"os"

	"sigs.k8s.io/yaml"
)

func main() {
	y, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to read from stdin: %v", err)
	}
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		log.Fatalf("unable to convert yaml into json: %v", err)
	}
	_, err = os.Stdout.Write(j)
	if err != nil {
		log.Fatalf("unable to write to stdout: %v", err)
	}
	os.Stdout.Write([]byte("\n"))
}
