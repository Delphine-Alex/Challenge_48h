package repository

import (
	"GoChallenge48/Go/models"
	"fmt"
)

// GetNewsletters get all newsletters from db
func GetAdmin() []models.Admin {
	rows, err := currentDB.Query("SELECT * FROM administrator")

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string

	var administrator []models.Admin

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password)

		if err != nil {
			panic(err)
		}

		fmt.Println(username)
		administrator = append(administrator, models.Admin{UUID: uuid, Username: username, Password: password})
	}

	return administrator
}