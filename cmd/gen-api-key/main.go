// Package main contains the main cli entry point
package main

import (
	"encoding/json/v2"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arizon-dread/gen-api-key/internal/models"
	"github.com/arizon-dread/gen-api-key/internal/services"
)

var (
	p = flag.String("p", "psk", "prefix")
	n = flag.String("n", "my-app", "application name. DO NOT use underscore (_) in the value")
	e = flag.String("e", "prod", "environment")
	j = flag.Bool("j", false, "output JSON")
	//h = flag.Bool("h", false, "output help text")
)

func main() {
	// Set help output
	flag.Usage = showHelp
	flag.Parse()

	if !validateFlagValues() {
		fmt.Printf("underscore (_) is not allowed in prefix, environment or app-name")
		os.Exit(2)
	}
	randStr, err := services.RandomHex()
	if err != nil {
		log.Fatalf("failed creating random hex string, %s", err)
	}
	res := models.Result{
		APIKey:  fmt.Sprintf("%s_%s_%s_%s", *p, *n, *e, randStr),
		Hash:    services.Hash(randStr),
		Env:     *e,
		AppName: *n,
	}
	if *j {
		output, err := json.Marshal(res)
		if err != nil {
			log.Printf("unable to marshal go struct into json")
			os.Exit(3)
		}
		fmt.Printf("%s", string(output))
		os.Exit(0)
	}
	fmt.Printf("APIKey: %s\nHash: %s", res.APIKey, res.Hash)
}

func showHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), `gen-api-key - a cli app for generating and hashing API keys
Usage %s [options]

options:
`, os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), ``)

	os.Exit(1)
}

func validateFlagValues() bool {
	if strings.Contains(*p, "_") {
		return false
	}
	if strings.Contains(*n, "_") {
		return false
	}
	if strings.Contains(*e, "_") {
		return false
	}
	return true
}
