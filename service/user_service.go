package service

import "mitton/gin-handson/model"

var users = []model.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

func GetUsers() []model.User {
	return users
}

func CreateUser(name string) model.User {
	var newID int

	if len(users) == 0 {
		newID = 1
	} else {
		newID = users[len(users)-1].ID + 1
	}

	user := model.User{
		ID:   newID,
		Name: name,
	}

	users = append(users, user)
	return user
}
