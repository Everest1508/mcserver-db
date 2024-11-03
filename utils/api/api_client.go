package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

type APIResponse struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *APIClient) Get(endpoint string, headers map[string]string) (*APIResponse, error) {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create GET request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(body),
	}, nil
}

func (c *APIClient) Post(endpoint string, payload interface{}, headers map[string]string) (*APIResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not create POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(body),
	}, nil
}

func (c *APIClient) Put(endpoint string, payload interface{}, headers map[string]string) (*APIResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshal payload: %w", err)
	}

	req, err := http.NewRequest("PUT", c.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not create PUT request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send PUT request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(body),
	}, nil
}

func (c *APIClient) Delete(endpoint string, headers map[string]string) (*APIResponse, error) {
	req, err := http.NewRequest("DELETE", c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create DELETE request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send DELETE request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(body),
	}, nil
}
