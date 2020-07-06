package models

type User struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Emails []string `json:emails`
}
