package models

import (
	"time"

	"eventapi.com/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserId      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	q := `INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	DateTime := time.Now()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetEvent(id int64) (*Event, error) {
	q := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(q, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func GetAllEvents() ([]Event, error) {
	q := `SELECT * FROM events`
	rows, err := db.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (event Event) Delete() error {
	query := `DELETE from events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	print(event.ID)
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (event Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ? , location = ?, dateTime = ?
		WHERE id = ? 
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}
