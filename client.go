package zendesk

import (
	"fmt"

	"gopkg.in/yaml.v2"
	"gopkg.in/resty.v0"
	"golang.org/x/net/context"
	"io/ioutil"
)

type Client struct {
	domain string
	client *resty.Client
	apiVersion string
}

func (c *Client) ZendeskApi() *Api {
	return &Api{
		client:  c,
		context: context.Background(),
	}
}

func (c *Client) toFullUrl(path string) string {
	return fmt.Sprintf("https://%v.zendesk.com/api/%s/%s", c.domain, c.apiVersion, path)
}

func (c *Client) get(path string, params map[string]string) (*resty.Response, error) {
	return c.client.R().SetQueryParams(params).Get(c.toFullUrl(path))
}

func (c *Client) post(path string, params interface{}, v interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(params).SetResult(v).Post(c.toFullUrl(path))
}

func (c *Client) put(path string, params interface{}, v interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(params).SetResult(v).Put(c.toFullUrl(path))
}

func (c *Client) delete(path string) (*resty.Response, error) {
	return c.client.R().Delete(c.toFullUrl(path))
}

func FromToken(config ZendeskConfiguration) *Client {
	username := fmt.Sprintf("%s/token", config.Email)

	client := resty.SetBasicAuth(username, config.Token)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")

	return &Client{
		domain: config.Domain,
		client: client,
		apiVersion: config.ApiVersion,
	}
}

func LoadConfiguration(path string) (ZendeskConfiguration) {
	zendeskConfiguration := ZendeskConfiguration{}
	zendeskConfigurationFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	err = yaml.Unmarshal(zendeskConfigurationFile, &zendeskConfiguration)

	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	return zendeskConfiguration
}