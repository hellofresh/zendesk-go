package zendesk

import "github.com/go-resty/resty/v2"

type PhoneNumberApiHandler struct {
	client Client
}

type MultiplePhoneNumber struct {
	Response []PhoneNumber `json:"phone_numbers"`
}

func (p PhoneNumberApiHandler) GetAll() ([]PhoneNumber, error) {
	response, err := p.client.get(
		"/channels/voice/phone_numbers.json",
		map[string]string{"minimal_mode": "true"},
	)

	if err != nil {
		return nil, err
	}

	return p.parseMultiObjects(response), err
}

func (p PhoneNumberApiHandler) parseMultiObjects(response *resty.Response) []PhoneNumber {
	var object MultiplePhoneNumber

	p.client.parseResponseToInterface(response, &object)

	return object.Response
}
