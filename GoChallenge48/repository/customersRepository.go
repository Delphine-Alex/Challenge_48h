package repository

import (
	"GoChallenge48/Go/models"
	"database/sql"
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
		err = rows.Scan(&uuid, &firstname, &lastname, &email, &username, &password)

		if err != nil {
			panic(err)
		}

		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String})
	}

	return users
}

// GetUserById return a unique user from db using id
func GetUserById(id string) models.User {
	rows, err := currentDB.Query(`SELECT * FROM "user" WHERE uuid = $1`, id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString
	var image sql.NullString
	var uuid_wishlist sql.NullString

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname, &image, &uuid_wishlist)

		if err != nil {
			panic(err)
		}
	}

	return models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String}
}
