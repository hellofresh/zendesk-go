package zendesk

import (
	"fmt"
	"gopkg.in/resty.v0"
)

type CommentApiHandler struct {
	client Client
}

type MultipleComment struct {
	Response []Comment `json:"comments"`
}

func (c CommentApiHandler) GetForTicket(ticketId int) ([]Comment, error) {
	response, err := c.client.get(
		fmt.Sprintf("/tickets/%d/comments.json", ticketId),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return c.parseMultiObjects(response), err
}

func (c CommentApiHandler) parseMultiObjects(response *resty.Response) []Comment {
	var object MultipleComment

	c.client.parseResponseToInterface(response, &object)

	return object.Response
}
