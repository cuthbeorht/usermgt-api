package service

import (
	"testing"
)

func TestCreateNewUserExpectUserWithEmailAndId(t *testing.T) {

	newUser := CreateNewUser{Email: "test.user@server.com"}

	createdUser, err := newUser.create()

	if createdUser.Email != "test.user@server.com" {
		t.Fatalf("Expected email to be %s but was %s", "test.user@server.com", createdUser.Email)
	}

	if createdUser.Id != 999 {
		t.Fatalf("Expected Id to be %d but was %d", 999, createdUser.Id)
	}

	if err != nil {
		t.Fatal("Error not exected: ", err)
	}
}

func TestGivenEmptyNewUserExcpectError(t *testing.T) {
	newUser := CreateNewUser{}

	createdUser, err := newUser.create()

	if err == nil {
		t.Fatalf("Expecting error when creating new user. Got New User with email %s and id %d", createdUser.Email, createdUser.Id)
	}

	if err.Error() != "new user is empty" {
		t.Fatalf("Expecting error 'New user is empty'. Got %s", err.Error())
	}

}
