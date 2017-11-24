package gogs

import (
	"fmt"
	"strings"
	"testing"

	gogs "github.com/gogits/go-gogs-client"
)

func TestGogsUserDelete(t *testing.T) {
	client := gogs.NewClient("http://localhost:3000", "2a4a2f289bd334b943013affdb353a1de8687870")
	if err := client.AdminDeleteUser("random"); err != nil {
		if strings.Contains(err.Error(), "404") {
			t.Log("User already deleted")
		} else {
			t.Fatal(err)
		}
	}
}

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

func TestGogsUserEdit(t *testing.T) {
	client := gogs.NewClient("http://localhost:3000", "2a4a2f289bd334b943013affdb353a1de8687870")
	disabled := false
	err := client.AdminEditUser("random", gogs.EditUserOption{
		Active: &disabled,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGogsUserRead(t *testing.T) {
	client := gogs.NewClient("http://localhost:3000", "2a4a2f289bd334b943013affdb353a1de8687870")
	user, err := client.GetUserInfo("random")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", user)
}
