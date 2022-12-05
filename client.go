package go_wix_store

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type Client struct {
	handler *resty.Client
	Product ProductService
}

type Option func(shopClient *Client)

func NewClientWithToken(apiToken string, storeId string) *Client {

	if apiToken == "" || storeId == "" {
		log.Fatalln("Wix Api Key and Store Id required")
	}

	restClient := newRestyClient(apiToken, storeId)

	return GenClient(WithRestyClient(restClient))
}

func WithRestyClient(r *resty.Client) Option {
	return func(c *Client) {
		c.handler = r
	}
}

func GenClient(opts ...Option) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	c.Product = &ProductServiceOp{client: c}

	return c
}

func newRestyClient(apiToken string, storeId string) *resty.Client {
	baseUrl := "https://www.wixapis.com/stores/v1/products"
	client := resty.New()

	client.SetBaseURL(baseUrl)
	client.SetHeader("Authorization", apiToken)
	client.SetHeader("wix-site-id", storeId)
	client.SetHeader("Content-Type", "application/json")

	return client
}
