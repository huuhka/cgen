package cgen

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"net/http"
	"time"
)

type OpenAiClient struct {
	config *Config
}

// CreateRequest creates a request for the OpenAI API
func (c OpenAiClient) CreateRequest(prompt string) (*http.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	body := c.generateRequestBody(prompt)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.config.Url, bytes.NewBuffer(body))

	if c.config.ApiKey != "" {
		req.Header.Set("api-key", c.config.ApiKey)
	} else {
		accessToken, err := c.config.AzureCredential.GetToken(ctx, policy.TokenRequestOptions{
			Scopes: []string{"https://openai.azure.com/.default"},
		})
		if err != nil {
			return nil, fmt.Errorf("unable to get access token: %w", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.Token))
	}

	return req, err
}

func (c OpenAiClient) generateRequestBody(prompt string) []byte {
	return nil
}