package db

import (
	"errors"
	"go-rest-test/models"
)

var (
	ErrNotFound error = errors.New("Not Found")
)

var users []models.User

func InitInMemoryDB() error {
	user1 := models.User{
		ID:   1,
		Name: "Kim",
		Age:  28,
		Emails: []string{
			"kim@korea.com",
			"kim@korea.co.kr",
		},
	}
	user2 := models.User{
		ID:   2,
		Name: "Lee",
		Age:  30,
		Emails: []string{
			"lee@korea.com",
			"lee@korea.co.kr",
		},
	}

	users = append(users, user1, user2)
	return nil
}

func GetAllUser() ([]models.User, error) {
	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, ErrNotFound
}

func UpdateUserByID(u models.User) error {
	for i := 0; i < len(users); i++ {
		if users[i].ID == u.ID {
			users[i].Age = u.Age
			users[i].Name = u.Name
			users[i].Emails = u.Emails
			return nil
		}
	}
	return ErrNotFound
}

func DeleteUserByID(id int) error {
	delIdx := -1
	for i, u := range users {
		if u.ID == id {
			delIdx = i
			break
		}
	}
	if delIdx < 0 {
		return ErrNotFound
	}
	// No matter if order changed
	users[delIdx] = users[len(users)-1]
	users[len(users)-1] = models.User{}
	users = users[:len(users)-1]
	return nil
}

func InsertUser(u models.User) error {
	for _, d := range users {
		if d.ID == u.ID {
			return errors.New("duplicate primary key: id")
		}
	}
	users = append(users, u)
	return nil
}
