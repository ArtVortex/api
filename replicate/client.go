package replicate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewClient() *Client {
	return &Client{
		// TODO: Create a new token, this one is for test only. And serve token as a secret.
		Authorization: "Token r8_2w5UERtphyGPx9iz2HLTTbvjRNEznwZ2YaBwl",
		API:           "https://api.replicate.com/v1/predictions",
	}
}

// Create a replicate prediction
// TODO: Add responseBody struct
func (c *Client) Create(request Request) (responseBody map[string]interface{}, err error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.API, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	return responseBody, err
}

// Get a replicate prediction
// TODO: Add responseBody struct
func (c *Client) Get(predictionId string) (responseBody map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", c.API+"/"+predictionId, nil)
	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	return responseBody, err
}

// list predictions
func (c *Client) ListPredictions() (responseBody map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", c.API, nil)
	req.Header.Add("Authorization", c.Authorization)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	return responseBody, err
}
