package dbrepo

import (
	"bnb-booking/internal/models"
	"errors"
	"time"
)

func (m *testDbRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDbRepo) InsertReservation(res models.Reservation) (int, error) {
	// if room id is 2, then fail; otherwise pass
	if res.RoomID == 2 {
		return 0, errors.New("Some error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDbRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("Some error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true is availability exists, and false id no availability exists
func (m *testDbRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDbRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *testDbRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}
