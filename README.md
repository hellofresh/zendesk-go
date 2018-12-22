# Zendesk API

This API is ready to send request for `Users` and `Tickets`

How to use?

## Configuration

Update the configuration.yml file inside config folder:

```
api_version: "v2"
domain: DOMAIN
email: EMAIL
token: TOKEN
rate_limit: REQUESTS_PER_MINUTE
```

The service uses API token to communicate with Zendesk.

## Connecting with Zendesk API
```
import "github.com/hellofresh/zendesk-go"

func main() {
    client, _ := zendesk.FromToken(
        zendesk.LoadConfiguration("./config/configuration.yml"),
    )
}
```

## Basic
```
user := zendesk.User {
    ID: [id],
    Name: [name],
    Email: [email],
}

client, _ := zendesk.FromToken(
    zendesk.LoadConfiguration("./config/configuration.yml"),
)

u, err := client.User().CreateOrUpdateUser(user)

if err != nil {
    log.Println(err)
} else {
    log.Println(u)
}
```

## User functions available

#### GetUser
Return the user

```
client.User().GetUser(1)
```

#### GetUsers
Return the list of all users

```
client.User().GetZendeskApi()
```

#### GetUsersByGroup
Return the list of all users in a group

```
client.User().GetUsersByGroup(4)
```

#### CreateOrUpdateUser
Create or update a user

```
client.User().CreateOrUpdateUser(user)
```

#### CreateUser
Create a new user

```
client.User().CreateUser(user)
```

#### UpdateUser
Update an existing user

```
client.User().UpdateUser(user)
```

#### DeleteUser
Delete an existing user

```
client.User().DeleteUser(1)
```

## Ticket function available

#### GetTicket
Return the ticket

```
client.Ticket().GetTicket(1)
```

#### GetTickets
Return the list of all tickets

```
client.Ticket().GetTickets()
```

#### GetRecentTickets
Return the recent tickets

```
client.Ticket().GetRecentTickets()
```

#### GetTicketsFromOrganization
Return the tickets from an organization

```
client.Ticket().GetTicketsFromOrganization(10)
```

#### GetManyTickets
Return a list of tickets

```
client.Ticket().GetManyTickets(int[]{1, 2, 3, 5})
```

#### GetRequestedTicketsFromUser
Return the tickets requests by an user

```
client.Ticket().GetRequestedTicketsFromUser(2)
```

#### GetCcdTicketsFromUser
```
client.Ticket().GetCcdTicketsFromUser(2)
```

#### GetAssignedTicketsFromUser
Return the tickets assigned from user

```
client.Ticket().GetAssignedTicketsFromUser(2)
```

#### CreateTicket
Create a new ticket

```
client.Ticket().CreateTicket(ticket)
```

#### UpdateTicket
Update a ticket

```
client.Ticket().UpdateTicket(ticket)
```

#### UpdateTicketMarkAsSpam
Update a ticket to mark as spam

```
client.Ticket().UpdateTicketMarkAsSpam(1)
```

#### DeleteTicket
Delete an existing ticket

```
client.Ticket().DeleteTicket(1)
```

#### LoadTicketComments
Load all comments of an existing ticket

```
client.Ticket().LoadTickectComments(&ticket)
```

#### LoadComments
Load all comments of many existing tickets

```
client.Ticket().LoadTickectComments(tickets)
```