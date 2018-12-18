package zendesk

import (
	"encoding/json"
	"fmt"
	"strings"

	resty "gopkg.in/resty.v0"
)

type Client struct {
	domain     string
	client     *resty.Client
	apiVersion string
}

func (c Client) User() UserApiHandler {
	return UserApiHandler{c}
}

func (c Client) Ticket() TicketApiHandler {
	return TicketApiHandler{c}
}

func (c Client) Search() SearchApiHandler {
	return SearchApiHandler{c}
}

func (c Client) TicketMetric() TicketMetricApiHandler {
	return TicketMetricApiHandler{c}
}

func (c Client) SatisfactionRating() SatisfactionRatingApiHandler {
	return SatisfactionRatingApiHandler{c}
}

func (c Client) toFullUrl(path string) string {
	return fmt.Sprintf("https://%v.zendesk.com/api/%s/%s", c.domain, c.apiVersion, path)
}

func (c Client) toPath(path string) string {
	baseURL := fmt.Sprintf("https://%v.zendesk.com/api/%s", c.domain, c.apiVersion)
	return strings.Replace(path, baseURL, "", -1)
}

func (c Client) get(path string, params map[string]string) (*resty.Response, error) {
	resp, err := c.client.R().SetQueryParams(params).Get(c.toFullUrl(path))

	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c Client) post(path string, params interface{}, v interface{}) (interface{}, error) {
	_, err := c.client.R().SetBody(params).SetResult(&v).Post(c.toFullUrl(path))

	if err != nil {
		return v, err
	}

	return v, nil
}

func (c Client) put(path string, params interface{}, v interface{}) (interface{}, error) {
	_, err := c.client.R().SetBody(params).SetResult(&v).Put(c.toFullUrl(path))

	if err != nil {
		return v, err
	}

	return v, nil
}

func (c Client) delete(path string) (*resty.Response, error) {
	return c.client.R().Delete(c.toFullUrl(path))
}

func (c Client) parseResponseToInterface(response *resty.Response, v interface{}) {
	json.Unmarshal(response.Body(), &v)
}
