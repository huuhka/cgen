package main

import (
	"cgen"
	"log"
	"os"
)

func main() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	endpoint := os.Getenv("OPENAI_ENDPOINT")
	if endpoint == "" {
		log.Fatal("OPENAI_ENDPOINT is not set")
	}

	deploymentName := os.Getenv("OPENAI_DEPLOYMENTNAME")
	if deploymentName == "" {
		log.Fatal("OPENAI_DEPLOYMENTNAME is not set")
	}

	config, err := cgen.NewConfig(endpoint, deploymentName, cgen.WithApiKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", *config)
}