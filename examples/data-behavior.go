package examples

import "fmt"

// Client is the data structure for an api client
type Client struct {
	apiKey string
	baseURL string
}

// NewClient creates a new client instance
func NewClient() Client {
	return Client{
		apiKey: "secret",
		baseURL: "https://vg.no",
	}
}

// GetPageInfo is a function attached to the client data structure
func (client Client) GetPageInfo() {
	fmt.Printf("%s\n", client.apiKey)
	fmt.Printf("%s\n", client.baseURL)
}