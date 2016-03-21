package zendesk
//
//import (
//	"fmt"
//	"net/url"
//
//	"golang.org/x/net/context"
//)
//
//type Via struct {}
//
//type Ticket struct {
//	Id int `json:"id,omitempty"`
//	Url string `json:"url,omitempty"`
//	ExternalId string `json:"external_id,omitempty"`
//	Type string `json:"type,omitempty"`
//	Subject string `json:"subject,omitempty"`
//	RawSubject string `json:"raw_subject,omitempty"`
//	Description string `json:"description,omitempty"`
//	Priority string `json:"priority,omitempty"`
//	Status string `json:"status,omitempty"`
//	Recipient string `json:"recipient,omitempty"`
//	RequesterId int `json:"requester_id,omitempty"`
//	SubmitterId int `json:"submitter_id,omitempty"`
//	AssigneeId int `json:"assignee_id,omitempty"`
//	OrganizationId int `json:"organization_id,omitempty"`
//	GroupId int `json:"group_id,omitempty"`
//	CollaboratorsId []int `json:"collaborators_id,omitempty"`
//	ForumTopicId int `json:"forum_topic_id,omitempty"`
//	ProblemId int `json:"problem_id,omitempty"`
//	HasIncidents bool `json:"has_incidents,omitempty"`
//	DueAt string `json:"due_at,omitempty"`
//	Tags []string `json:"tags,omitempty"`
//	Via *Via `json:"via,omitempty"`
//	CustomFields []string `json:"custom_fields,omitempty"`
//	SatisfactionRating []string `json:"satisfaction_rating,omitempty"`
//	SharingAgreementIds []int `json:"sharing_agreement_ids,omitempty"`
//	FollowupIds []int `json:"followup_ids,omitempty"`
//	TicketFormId int `json:"ticket_form_id,omitempty"`
//	BrandId int `json:"brand_id,omitempty"`
//	CreatedAt string `json:"created_at,omitempty"`
//	UpdatedAt string `json:"updated_at,omitempty"`
//}
//
//type TicketApi struct {
//	client  *Client
//	context context.Context
//}
//
//// Make the http request GET
//func (api *TicketApi) ticketGet(path string, params *url.Values) (Ticket, error) {
//	response := struct {
//		ticket Ticket `json:"ticket"`
//	}{}
//
//	err := api.client.get(api.context, path, params, &response)
//
//	if err != nil {
//		return Ticket{}, err
//	}
//
//	return response.ticket, nil
//}
//
//// Make the http request POST
//func (api *TicketApi) ticketPost(path string, payload interface{}) (Ticket, error) {
//	response := struct {
//		ticket Ticket `json:"ticket"`
//	}{}
//
//	err := api.client.post(api.context, path, payload, &response)
//
//	if err != nil {
//		return Ticket{}, err
//	}
//
//	return response.ticket, nil
//}
//
//// Make the http request DELETE
//func (api *TicketApi) ticketDelete(path string) (Ticket, error) {
//	response := struct {
//		ticket Ticket `json:"ticket"`
//	}{}
//
//	err := api.client.delete(api.context, path, &response)
//
//	if err != nil {
//		return Ticket{}, err
//	}
//
//	return response.ticket, nil
//}
//
//// Make the http request PUT
//func (api *TicketApi) ticketUpdate(path string, payload interface{}) (Ticket, error) {
//	response := struct {
//		ticket Ticket `json:"ticket"`
//	}{}
//
//	err := api.client.put(api.context, path, payload, &response)
//
//	if err != nil {
//		return Ticket{}, err
//	}
//
//	return response.ticket, nil
//}
//
//// Show function
//// search and show for a Ticket by id
//func (api *TicketApi) showTicket(id int) (Ticket, error) {
//	path := fmt.Sprintf("/api/v2/tickets/%d.json", id)
//
//	return api.ticketGet(path, nil)
//}
//
//// Create function
//// create a new Ticket in zendesk
//func (api *TicketApi) createTicket(ticket Ticket) (Ticket, error) {
//	return api.ticketPost("/api/v2/tickets.json", map[string]Ticket{"ticket": ticket})
//}
//
//// Update function
//// update an existing Ticket in zendesk
//func (api *TicketApi) updateTicket(ticket Ticket, id int) (Ticket, error) {
//	path := fmt.Sprintf("/api/v2/tickets/%d.json", id)
//
//	return api.ticketUpdate(path, map[string]Ticket{"ticket": ticket})
//}
//
//// Delete
//// delete the Ticket by id
//func (api *TicketApi) deleteTicket(id int) (Ticket, error) {
//	path := fmt.Sprintf("/api/v2/tickets/%d.json", id)
//
//	return api.ticketDelete(path)
//}