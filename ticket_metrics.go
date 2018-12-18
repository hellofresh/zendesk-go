package zendesk

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

type TicketMetricApiHandler struct {
	client Client
}

func (s TicketMetricApiHandler) TicketMetrics() (TicketMetrics, error) {
	response, err := s.client.get(
		fmt.Sprintf("/ticket_metrics.json"),
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s TicketMetricApiHandler) TicketMetricsPage(page int) (TicketMetrics, error) {
	response, err := s.client.get(
		fmt.Sprintf("/ticket_metrics.json?page=%d", page),
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s TicketMetricApiHandler) TicketMetricsByName(metricName string) (TicketMetrics, error) {
	response, err := s.client.get(
		fmt.Sprintf("/ticket_metrics/%s.json", metricName),
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s TicketMetricApiHandler) parseResults(response *resty.Response) TicketMetrics {
	var object TicketMetrics

	s.client.parseResponseToInterface(response, &object)

	return object
}
