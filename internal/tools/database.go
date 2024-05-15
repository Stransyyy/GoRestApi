package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type TaskDetails struct {
	Tasks    int //number of tasks of that person
	Username string
}

type DatabaseInterface interface {
	GetUserLogInDetails(username string) *LoginDetails
	GetTaskDetails(username string) *TaskDetails
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
