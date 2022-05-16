package repository

import (
	"database/sql"
	"GoChallenge48/Go/models"
)

// GetUsers return a user from db

func GetCustomers() []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "customers" ORDER BY lastname ASC`)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString

	var users []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String})
	}

	return users
}