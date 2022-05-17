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
func GetCustomerById(id string) models.User {
	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE uuid = $1`, id)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname sql.NullString
	var lastname sql.NullString

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

		if err != nil {
			panic(err)
		}
	}

	return models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String}
}


// GetUserByEmail return a user from db using email
func GetCustomerByEmail(mail string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE email = $1`, mail)

	if err != nil {
		panic(err)
	}

	var uuid string
	var username string
	var password string
	var email string
	var firstname string
	var lastname string

	var Customers []models.User

	for rows.Next() {
		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

		if err != nil {
			panic(err)
		}

		Customers = append(Customers, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname, LastName: lastname})
	}

	return Customers
}

// GetUserByUsername return a user from db using username
func GetCustomerByUsername(name string) []models.User {
	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE username = $1`, name)

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

// PostUser create a new user in db
func PostUser(newUser models.CreateUser) {
	// dynamic
	insertDynStmt := `insert into "customers"("username", "password", "email") values($1, $2, $3)`

	_, err := currentDB.Exec(insertDynStmt, newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		panic(err)
	}
}