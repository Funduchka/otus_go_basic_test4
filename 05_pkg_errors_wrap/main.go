package main

import (
	"errors"
	"fmt"
)

const testEmail = "test@mail.ru"

func getUserEmail(id int) (string, error) {
	if id = 1 {return testEmail, nil}

	return "", errors.New("email not found")
}

func loginUserByID(id int) (bool, error) {
	email, err := getUserEmail(id)
	if err != nil { return false, errors.Wrap(err, "unable to get email") }
	if email != testEmail { return false, errors.New("unable to login user") }

	return true, nil
}

func main() {
	if _, err := loginUserByID(10); err != nil {
		fmt.Println("%+v\n", err)
	}
	else {
		fmt.Println("success")
	}
}