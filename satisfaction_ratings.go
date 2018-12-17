package zendesk

import (
	"fmt"
	"time"

	resty "gopkg.in/resty.v0"
)

type SatisfactionRatingApiHandler struct {
	client Client
}

func (s SatisfactionRatingApiHandler) SatisfactionRatings(score string, startDate *time.Time, stopDate *time.Time) (SatisfactionRatings, error) {
	var urlString string

	if len(score) > 0) {
		formatString =  "/satisfaction_ratings.json?start_time=%d&end_time=%d&score=%s"
		urlString :=		fmt.Sprintf(formatString, startDate.Unix(), stopDate.Unix(), score),
	} else {
		formatString:= "/satisfaction_ratings.json?start_time=%d&end_time=%d"
		urlString :=		fmt.Sprintf(formatString, startDate.Unix(), stopDate.Unix()),
	}
	
	response, err := s.client.get(
		urlString
		nil,
	)

	if err != nil {

	}

	return s.parseResults(response), err
}

func (s SatisfactionRatingApiHandler) parseResults(response *resty.Response) SatisfactionRatings {
	var object SatisfactionRatings

	s.client.parseResponseToInterface(response, &object)

	return object
}
