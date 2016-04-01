package zendesk

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	client, _ := FromEnv();

	user, _ := client.Users().GetUser(211722549);

	if user.Id != 211722549 {
		t.Fatal("ERROR")
	}
}