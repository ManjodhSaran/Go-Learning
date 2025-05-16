package models

import (
	"eventapi.com/db"
)

type EventRegistration struct {
	ID      int64 `json:"id"`
	EventId int64 `json:"event_id"`
	UserId  int64 `json:"user_id"`
}

type DetailedRegistration struct {
	RegisterId       int64  `json:"register_id"`
	EventId          int64  `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventLocation    string `json:"event_location"`
	UserId           int64  `json:"user_id"`
	UserEmail        string `json:"user_email"`
}

func (e *EventRegistration) Save() error {
	q := `INSERT INTO event_registrations(event_id,user_id)
	VALUES (?,?)`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.EventId, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func DeleteRegistration(id int64) error {
	q := `DELETE FROM event_registrations WHERE id = ?`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func GetRegisterById(id int64) (*EventRegistration, error) {
	q := `SELECT * FROM event_registrations WHERE id = ?`

	row := db.DB.QueryRow(q, id)

	var registration EventRegistration
	err := row.Scan(&registration.ID, &registration.EventId, &registration.UserId)

	if err != nil {
		return nil, err
	}
	return &registration, nil

}

func GetEventRegistrations(id int64) ([]EventRegistration, error) {
	q := `SELECT * FROM event_registrations WHERE event_id = ?`

	rows, err := db.DB.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registrations []EventRegistration

	for rows.Next() {
		var registration EventRegistration
		err := rows.Scan(&registration.ID, &registration.EventId, &registration.UserId)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func CheckIsRegistered(userId, eventId int64) (int64, error) {
	q := `SELECT id FROM event_registrations WHERE user_id = ? AND event_id = ?`

	row := db.DB.QueryRow(q, userId, eventId)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		// If no rows are found, return 0 and nil error
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}

func GetDetailedRegistration(id int64) (*DetailedRegistration, error) {
	// populate the event_id and user_id fields right join
	q := `SELECT er.id as register_id, er.event_id, er.user_id, 
	             e.name as event_name, e.description as event_description, e.location as event_location,
	            u.email as user_email  
	      FROM event_registrations er
	      LEFT JOIN events e ON er.event_id = e.id
	      LEFT JOIN users u ON er.user_id = u.id
	      WHERE er.event_id = ?`

	rows, err := db.DB.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var registrations []DetailedRegistration
	for rows.Next() {
		var registration DetailedRegistration
		err := rows.Scan(
			&registration.RegisterId,
			&registration.EventId,
			&registration.UserId,
			&registration.EventName,
			&registration.EventDescription,
			&registration.EventLocation,
			&registration.UserEmail,
		)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(registrations) == 0 {
		return nil, nil
	}
	return &registrations[0], nil
}
