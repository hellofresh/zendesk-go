package zendesk

import (
	"golang.org/x/net/context"
	"gopkg.in/resty.v0"
)

type Api struct {
	client  *Client
	context context.Context
}

func (api *Api) getHttpRequest(path string, params map[string]string, response interface{}) (interface{}, error) {
	u := response

	_, err := api.client.get(path, params, &u)

	if err != nil {
		return response, err
	}

	return u, nil
}

func (api *Api) postHttpRequest(path string, payload interface{}, response interface{}) (interface{}, error) {
	_, err := api.client.post(path, payload, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api *Api) deleteHttpRequest(path string) (*resty.Response, error) {
	response, err := api.client.delete(path)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api *Api) updateHttpRequest(path string, payload interface{}, response interface{}) (interface{}, error) {
	_, err := api.client.put(path, payload, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api *Api) updateHttpRequestGetResponse(path string, payload interface{}) (*resty.Response, error) {
	response, err := api.client.put(path, payload, nil)

	if err != nil {
		return nil, err
	}

	return response, nil
}