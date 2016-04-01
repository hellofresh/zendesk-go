package zendesk

import (
	"fmt"
	"os"

	"gopkg.in/resty.v0"
	"golang.org/x/net/context"
)

type Client struct {
	domain string
	client *resty.Client
}

func (c *Client) Users() *Api {
	return &Api{
		client:  c,
		context: context.Background(),
	}
}

func (c *Client) toFullUrl(path string) string {
	return fmt.Sprintf("https://%v.zendesk.com/api/%s%s", os.Getenv("ZENDESK_API_VERSION"), c.domain, path)
}

func (c *Client) get(path string, params map[string]string, v interface{}) (*resty.Response, error) {
	return c.client.R().SetQueryParams(params).SetResult(v).Get(c.toFullUrl(path))
}

func (c *Client) post(path string, params interface{}, v interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(params).SetResult(v).Post(c.toFullUrl(path))
}

func (c *Client) delete(path string) (*resty.Response, error) {
	return c.client.R().Delete(c.toFullUrl(path))
}

func (c *Client) put(path string, params interface{}, v interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(params).SetResult(v).Put(c.toFullUrl(path))
}

func FromEnv() (*Client, error) {
	domain := os.Getenv("ZENDESK_DOMAIN")
	email := os.Getenv("ZENDESK_EMAIL")
	token := os.Getenv("ZENDESK_TOKEN")

	return FromToken(domain, email, token), nil
}

func FromToken(domain, email string, token string) *Client {
	username := fmt.Sprintf("%s/token", email)

	client := resty.SetBasicAuth(username, token)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		domain: domain,
		client: client,
	}
}
