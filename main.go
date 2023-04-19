package main

import (
	"fmt"
	"go-xml-parser/internal/app"
	. "go-xml-parser/internal/models"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	f, err := os.Open("internal/config/config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable configuration file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration file currupted: %v\n", err)
		os.Exit(1)
	}
	app.Run(cfg)
}
