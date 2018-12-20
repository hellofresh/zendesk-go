package zendesk

import "gopkg.in/resty.v0"

type ZendeskApi interface {
	GetById(id int) (interface{}, error)
	GetAll() ([]interface{}, error)
	Create(v interface{}) (interface{}, error)
	CreateOrUpdate(v interface{}) (interface{}, error)
	CreateOrUpdateMany(v []interface{}) (interface{}, error)
	Update(v interface{}) (interface{}, error)
	Delete(id int) (int, error)

	// Parse response
	parseSingleObject(response *resty.Response) interface{}
	parseMultiObjects(response *resty.Response) []interface{}
}
