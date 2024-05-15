package tools

import (
	"log"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"admin": {
		AuthToken: "admin123",
		Username:  "admin",
	},
	"user": {
		AuthToken: "QWLPG",
		Username:  "user",
	},
	"alan": {
		AuthToken: "FSDLK",
		Username:  "alan",
	},
	"adam": {
		AuthToken: "ASD$L",
		Username:  "adam",
	},
}

var mockTaskDetails = map[string]TaskDetails{
	"admin": {
		Tasks:    10,
		Username: "admin",
	},
	"user": {
		Tasks:    5,
		Username: "user",
	},
	"alan": {
		Tasks:    3,
		Username: "alan",
	},
	"adam": {
		Tasks:    2,
		Username: "adam",
	},
}

func (db *mockDB) GetUserLogInDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		log.Println("User not found")
		return nil
	}

	return &clientData
}

func (db *mockDB) GetTaskDetails(username string) *TaskDetails {
	time.Sleep(1 * time.Second)

	var clientData = TaskDetails{}
	clientData, ok := mockTaskDetails[username]
	if !ok {
		log.Println("User not found")
		return nil
	}

	return &clientData
}

func (db *mockDB) SetUpDatabase() error {
	return nil
}
