package main

import (
	"cgen"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading input: %s", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	endpoint := os.Getenv("OPENAI_ENDPOINT")
	if endpoint == "" {
		log.Fatal("OPENAI_ENDPOINT is not set")
	}

	deploymentName := os.Getenv("OPENAI_DEPLOYMENT_NAME")
	if deploymentName == "" {
		log.Fatal("OPENAI_DEPLOYMENT_NAME is not set")
	}

	config, err := cgen.NewConfig(endpoint, deploymentName, cgen.WithApiKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	client, err := cgen.NewOpenAiClient(config)

	commitMsg, err := client.GetCommitMessage(string(inputBytes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", commitMsg)
}