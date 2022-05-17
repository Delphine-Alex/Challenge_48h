package repository

import (
	"GoChallenge48/Go/models"
	"database/sql"
)

// GetUsers return a user from db

func GetEvents() []models.Jointure {
	rows, err := currentDB.Query(`SELECT * FROM event JOIN organizer ON event.uuid_organizator = organizer.uuid`)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var _type string
	var location string
	var start string
	var end string
	var uuid_organizator sql.NullString

	var uuid_orga string
	var name string
	var referent string
	var email string
	var username string
	var password string

	var events []models.Jointure

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &_type, &location, &start, &end, &uuid_organizator, &uuid_orga, &name, &referent, &email, &username, &password)

		if err != nil {
			panic(err)
		}

		events = append(events, models.Jointure{Uuid: uuid, Title: title, TypeInfo: _type, Location: location, Start: start, End: end, Uuid_organizator: uuid_organizator.String, Uuid_orga: uuid_orga, Name: name, Referent: referent, Email: email, Username: username, Password: password})
	}

	return events
}

// GetUserByUsername return a user from db using username
func GetEventByName(name string) []models.Event {
	rows, err := currentDB.Query(`SELECT * FROM "event" WHERE title = $1`, name)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var _type string
	var location string
	var start string
	var end string
	var uuid_organizator sql.NullString;

	var events []models.Event

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &_type, &location, &start, &end, &uuid_organizator)

		if err != nil {
			panic(err)
		}

		events = append(events, models.Event{UUID: uuid, Title: title, Type: _type, Location: location, Start: start, End: end, UUID_organizator: uuid_organizator.String})
	}

	return events
}

// PostUser create a new user in db
func PostEvent(newUser models.CreateEvent) {
	// dynamic
	insertDynStmt := `insert into "event"("title", "type", "location","start","end", "uuid_organizator") values($1, $2, $3, $4, $5,$6)`

	_, err := currentDB.Exec(insertDynStmt, newUser.Title, newUser.Type, newUser.Location, newUser.Start, newUser.End, newUser.UUID_organizator)
	if err != nil {
		panic(err)
	}
}
