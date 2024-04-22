package db

import (
	"math/rand"
	"time"

	"github.com/Stransyyy/GoRestApi/models"
	name "github.com/random-names/go"
)

func UserList() ([]models.Users, error) {
	var users []models.Users

	_, err := randEmail()
	if err != nil {
		return []models.Users{}, err
	}

	// for _, v := range s {
	// 	users = append(users, models.Users{
	// 		ID:    ,
	// 		Email: v,
	// 	})
	// }

	return users, nil
}

func randEmail() ([]string, error) {
	name, err := name.GetRandomNames("census-90/male.first", &name.Options{
		Number: 100,
	})

	var names []string

	if err != nil {
		return names, err
	}

	for _, n := range name {
		names = append(names, n+"@gmail.com")
	}

	return names, nil
}

func init() {
	rand.New(time.Now().UnixNano())
}
