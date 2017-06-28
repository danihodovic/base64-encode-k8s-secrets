package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func usage() string {
	return fmt.Sprintf(
		`
Usage: %s <secret.yaml>
`, os.Args[0])
}

type secret struct {
	Data       map[string]string `yaml:"data"`
	ApiVersion string            `yaml:"apiVersion"`
	Kind       string            `yaml:"kind"`
	Type       string            `yaml:"type"`
	Metadata   struct {
		CreationTimestamp string `yaml:"creationTimestamp"`
		Name              string `yaml:"name"`
		Namespace         string `yaml:"namespace"`
		ResourceVersion   string `yaml:"resourceVersion"`
		SelfLink          string `yaml:"selfLink"`
		Uid               string `yaml:"uid"`
	} `yaml:"metadata"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage())
		os.Exit(1)
	}

	fileBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	s := secret{}
	err = yaml.Unmarshal(fileBytes, &s)
	if err != nil {
		panic(err)
	}

	for key, value := range s.Data {
		s.Data[key] = base64.StdEncoding.EncodeToString([]byte(value))
	}

	newYaml, err := yaml.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(newYaml))
}
