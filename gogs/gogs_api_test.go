package gogs

import (
	"fmt"
	"testing"

	gogs "github.com/gogits/go-gogs-client"
)

func TestGogsUserCreate(t *testing.T) {
	client := gogs.NewClient("http://localhost:3000", "2a4a2f289bd334b943013affdb353a1de8687870")
	user, err := client.AdminCreateUser(gogs.CreateUserOption{
		Email:     "random@example.com",
		LoginName: "random",
		Username:  "random",
	})

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)

}
