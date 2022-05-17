package repository

import (
	"GoChallenge48/Go/models"
)

// GetUsers return a user from db

func GetOrganizer() []models.Organizer {
	rows, err := currentDB.Query(`SELECT * FROM "organizer" ORDER BY username ASC`)

	if err != nil {
		panic(err)
	}

	var uuid string
	var name string
	var referent string
	var email string
	var username string
	var password string

	var organizers []models.Organizer

	for rows.Next() {
		err = rows.Scan(&uuid, &name, &referent, &email, &username, &password)

		if err != nil {
			panic(err)
		}

		organizers = append(organizers, models.Organizer{UUID: uuid, Name: name, Referent: referent, Email: email, Username: username, Password: password})
	}

	return organizers
}