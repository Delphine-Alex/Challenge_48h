package repository

import (
	"GoChallenge48/Go/models"
)

// GetUsers return a user from db

func GetEvents() []models.Event {
	rows, err := currentDB.Query(`SELECT * FROM "event" ORDER BY start ASC`)

	if err != nil {
		panic(err)
	}

	var uuid string
	var title string
	var _type string
	var location string
	var start string
	var end string
	var uuid_wishlist string

	var events []models.Event

	for rows.Next() {
		err = rows.Scan(&uuid, &title, &_type, &location, &start, &end,&uuid_wishlist)

		if err != nil {
			panic(err)
		}

		events = append(events, models.Event{UUID: uuid, Title: title, Type: _type, Location: location, Start: start, End: end, UUID_wishlist: uuid_wishlist})
	}

	return events
}