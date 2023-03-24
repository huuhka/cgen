package cgen

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// CreateUrl creates a url for the request
func CreateUrl(endpoint, deploymentName string) string {
	apiVersion := "2023-03-15"
	return fmt.Sprintf("%s/openai/deployments/%s/chat/completions?api-version=%s", endpoint, deploymentName, apiVersion)
}

// Config is the configuration for calling the Azure OpenAI API
type Config struct {
	Url             string
	ApiKey          string
	SystemContext   Message
	AzureCredential *azidentity.DefaultAzureCredential
}

func NewConfig(endpoint, deploymentName string, opts ...ConfigOpt) (*Config, error) {
	systemRole := "You generate Git commit messages based on provided output from git diff commands. " +
		"Keep the replies short and to the point. Do not include any explanation, only the commit message."

	c := Config{
		Url:           CreateUrl(endpoint, deploymentName),
		SystemContext: NewSystemMessage(systemRole),
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return nil, err
		}
	}

	if c.AzureCredential == nil && c.ApiKey == "" {
		return nil, fmt.Errorf("no authentication method provided, please provide an API key or Azure credential")
	}

	return &c, nil
}

type ConfigOpt func(*Config) error

func WithApiKey(apiKey string) ConfigOpt {
	return func(c *Config) error {
		c.ApiKey = apiKey
		return nil
	}
}

func WithAzureCredential() ConfigOpt {
	return func(c *Config) error {
		cred, err := LoginWithDefaultCredential()
		if err != nil {
			return err
		}

		c.AzureCredential = cred
		return nil
	}
}