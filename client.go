package zendesk

import (
	"fmt"

	"gopkg.in/resty.v0"
	"golang.org/x/net/context"
	"log"
)

type Client struct {
	domain string
	client *resty.Client
}
//
func (c *Client) Users() *UserApi {
	return &UserApi{
		client:  c,
		context: context.Background(),
	}
}

//func (c *Client) Ticket() *TicketApi {
//	return &TicketApi{
//		client:  c,
//		context: context.Background(),
//	}
//}

func (c *Client) toFullUrl(path string) string {
	return fmt.Sprintf("https://%v.zendesk.com%s", c.domain, path)
}

func (c *Client) get(path string, params map[string]string, v interface{}) (*resty.Response, error) {
	return c.client.R().SetQueryParams(params).SetResult(v).Get(c.toFullUrl(path))
}

func (c *Client) post(path string, params interface{}, v interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(params).SetResult(v).Post(c.toFullUrl(path))
}

func (c *Client) delete(path string, v interface{}) (*resty.Response, error) {
	return c.client.R().SetResult(v).Delete(c.toFullUrl(path))
}

func (c *Client) put(path string, params interface{}, v interface{}) (*resty.Response, error) {
	resp, _ := c.client.R().SetBody(params).SetResult(&v).Put(c.toFullUrl(path))

	log.Println(c.client.R().SetBody(params).SetResult(&v).Body)
	log.Println(v)
	log.Println("Response:", resp)

	return resp, nil
}

func FromEnv() (*Client, error) {
	domain := "hellofresh1"
	email := "ebe@hellofresh.com"
	token := "nqn0lixOpzcYNK14VthJiIy9kA9pqrxrn13wiEIX"

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