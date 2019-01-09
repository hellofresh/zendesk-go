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

u, err := client.User().CreateOrUpdate(user)

if err != nil {
    log.Println(err)
} else {
    log.Println(u)
}
```

## User functions available

#### Get user
Return the user

```
client.User().GetById(1)
```

#### Get users
Return the list of all users

```
client.User().GetAll()
```

#### Get agents
Return the list of all agents

```
client.User().GetAllAgents(4)
```

#### Create or update user
Create or update a user

```
client.User().CreateOrUpdate(user)
```

#### Create user
Create a new user

```
client.User().Create(user)
```

#### Update user
Update an existing user

```
client.User().Update(user)
```

#### Delete user
Delete an existing user

```
client.User().Delete(1)
```

## Ticket function available

#### Get ticket
Return the ticket

```
client.Ticket().GetById(1)
```

#### Get all tickets
Return the list of all tickets

```
client.Ticket().GetAll()
```

#### Create ticket
Create a new ticket

```
client.Ticket().Create(ticket)
```

#### Update ticket
Update a ticket

```
client.Ticket().Create(ticket)
```

#### Delete ticket
Delete an existing ticket

```
client.Ticket().Delete(1)
```