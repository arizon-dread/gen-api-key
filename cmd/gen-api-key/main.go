// Package main contains the main cli entry point
package main

import (
	"encoding/json/v2"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/arizon-dread/gen-api-key/internal/models"
	"github.com/arizon-dread/gen-api-key/internal/services"
)

var (
	p = flag.String("p", "psk", "Prefix. Set to \"\" to clear")
	n = flag.String("n", "", "Application name")
	e = flag.String("e", "", "Environment")
	j = flag.Bool("j", false, "Output JSON")
)

func main() {
	// Set help output
	flag.Usage = showHelp
	flag.Parse()

	if !services.ValidateFlagValues(*p, *n, *e) {
		fmt.Printf("underscore (_) is not allowed in prefix, environment or app-name")
		os.Exit(2)
	}
	randStr, err := services.RandomHex()
	if err != nil {
		log.Fatalf("failed creating random hex string, %s", err)
	}

	prefix := fmt.Sprintf("%s%s%s", services.AddSepIfValue(*p), services.AddSepIfValue(*n), services.AddSepIfValue(*e))
	apiKey := randStr

	if len(prefix) > 0 {
		apiKey = fmt.Sprintf("%s%s", prefix, apiKey)
	}

	res := models.Result{
		APIKey:  apiKey,
		Hash:    services.Hash(randStr),
		Env:     *e,
		AppName: *n,
		Prefix:  *p,
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
	os.Exit(0)
}

func showHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), `gen-api-key - A cli app for generating and hashing API keys
Usage %s [options]

Note: Do not use underscores in the flag values!

options:
`, os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), ``)

	os.Exit(1)
}
