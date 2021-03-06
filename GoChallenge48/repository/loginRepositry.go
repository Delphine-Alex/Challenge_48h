package repository

// import (
// 	"database/sql"
// 	"GoChallenge48/Go/models"
// )

// // //GetUsers return a user from db

// func GetUsers() []models.User {
// 	rows, err := currentDB.Query(`SELECT * FROM "customers" ORDER BY lastname ASC`)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var uuid string
// 	var password string
// 	var email string
	

// 	var users []models.User

// 	for rows.Next() {
// 		err = rows.Scan(&uuid,&password, &email)

// 		if err != nil {
// 			panic(err)
// 		}

// 		users = append(users, models.User{UUID: uuid, Password: password, Email: email})
// 	}

// 	return users
// }

// // GetUserById return a unique user from db using id
// func GetLoginById(id string) models.User {
// 	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE uuid = $1`, id)

// 	if err != nil {
// 		panic(err)
// 	}

	
// 	var password string
// 	var email string
	

// 	for rows.Next() {
// 		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	return models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String}
// }

// //////////////////////////////////////
// ///			POST
// //////////////////////////////////////

// // GetUserByEmail return a user from db using email
// func GetLoginByEmail(mail string) []models.User {
// 	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE email = $1`, mail)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var uuid string
// 	var username string
// 	var password string
// 	var email string
// 	var firstname sql.NullString
// 	var lastname sql.NullString

// 	var Customers []models.User

// 	for rows.Next() {
// 		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

// 		if err != nil {
// 			panic(err)
// 		}

// 		Customers = append(Customers, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String})
// 	}

// 	return Customers
// }

// // GetUserByUsername return a user from db using username
// func GetLoginByUsername(name string) []models.User {
// 	rows, err := currentDB.Query(`SELECT * FROM "customers" WHERE username = $1`, name)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var uuid string
// 	var username string
// 	var password string
// 	var email string
// 	var firstname sql.NullString
// 	var lastname sql.NullString

// 	var users []models.User

// 	for rows.Next() {
// 		err = rows.Scan(&uuid, &username, &password, &email, &firstname, &lastname)

// 		if err != nil {
// 			panic(err)
// 		}

// 		users = append(users, models.User{UUID: uuid, Username: username, Password: password, Email: email, FirstName: firstname.String, LastName: lastname.String})
// 	}

// 	return users
// }

// // PostUser create a new user in db
// func PostLogin(newUser models.CreateUser) {
// 	// dynamic
// 	insertDynStmt := `insert into "customers"("username", "password", "email") values($1, $2, $3)`

// 	_, err := currentDB.Exec(insertDynStmt, newUser.Username, newUser.Password, newUser.Email)
// 	if err != nil {
// 		panic(err)
// 	}
// }


