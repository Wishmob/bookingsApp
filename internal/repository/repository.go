package repository

import "github.com/Wishmob/bookingsApp/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
