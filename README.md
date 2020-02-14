# Zendesk API

This API is ready to send request for `Users` and `Tickets`

How to use?

## Configuration

Update the configuration.yml file inside config folder:

```
domain: DOMAIN
email: EMAIL
password: PASSWORD
token: TOKEN
```

The service uses API token to communicate with Zendesk.

## Connecting with Zendesk API
```
import "github.com/hellofresh/zendesk-go"

func main() {
    client, _ := zendesk.FromEnv(
        zendesk.LoadConfiguration("./config/configuration.yml"),
    )
}
```

## Basic
```
user := zendesk.User {
    Id: [id],
    Name: [name],
    Email: [email],
}

client, _ := zendesk.FromEnv(
    zendesk.LoadConfiguration("./config/configuration.yml"),
)

u, err := client.ZendeskApi().CreateOrUpdateUser(user)

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
client.ZendeskApi().GetUser(1)
```

#### GetUsers
Return the list of all users

```
client.ZendeskApi().GetZendeskApi()
```

#### GetUsersByGroup
Return the list of all users in a group

```
client.ZendeskApi().GetUsersByGroup(4)
```

#### CreateOrUpdateUser
Create or update a user

```
client.ZendeskApi().CreateOrUpdateUser(user)
```

#### CreateUser
Create a new user

```
client.ZendeskApi().CreateUser(user)
```

#### UpdateUser
Update an existing user

```
client.ZendeskApi().UpdateUser(user)
```

#### DeleteUser
Delete an existing user

```
client.ZendeskApi().DeleteUser(1)
```

## Ticket function available

#### GetTicket
Return the ticket

```
client.ZendeskApi().GetTicket(1)
```

#### GetTickets
Return the list of all tickets

```
client.ZendeskApi().GetTickets()
```

#### GetRecentTickets
Return the recent tickets

```
client.ZendeskApi().GetRecentTickets()
```

#### GetTicketsFromOrganization
Return the tickets from an organization

```
client.ZendeskApi().GetTicketsFromOrganization(10)
```

#### GetManyTickets
Return a list of tickets

```
client.ZendeskApi().GetManyTickets(int[]{1, 2, 3, 5})
```

#### GetRequestedTicketsFromUser
Return the tickets requests by an user

```
client.ZendeskApi().GetRequestedTicketsFromUser(2)
```

#### GetCcdTicketsFromUser
```
client.ZendeskApi().GetCcdTicketsFromUser(2)
```

#### GetAssignedTicketsFromUser
Return the tickets assigned from user

```
client.ZendeskApi().GetAssignedTicketsFromUser(2)
```

#### CreateTicket
Create a new ticket

```
client.ZendeskApi().CreateTicket(ticket)
```

#### UpdateTicket
Update a ticket

```
client.ZendeskApi().UpdateTicket(ticket)
```

#### UpdateTicketMarkAsSpam
Update a ticket to mark as spam

```
client.ZendeskApi().UpdateTicketMarkAsSpam(1)
```

#### DeleteTicket
Delete an existing ticket

```
client.ZendeskApi().DeleteTicket(1)
```

License
-------

    Copyright (C) 2016 HelloFresh SE

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
