package zendesk

import (
	"golang.org/x/net/context"
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

func (api *Api) deleteHttpRequest(path string, response interface{}) (interface{}, error) {
	_, err := api.client.delete(path, &response)

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
