package main

import (
	"github.com/Wishmob/bookingsApp/internal/config"
	"github.com/Wishmob/bookingsApp/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals_quarters", handlers.Repo.Generals)
	mux.Get("/majors_suit", handlers.Repo.Majors)
	mux.Get("/make_reservation", handlers.Repo.Reservation)
	mux.Post("/make_reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation_summary", handlers.Repo.ReservationSummary)
	mux.Get("/search_availability", handlers.Repo.Availability)
	mux.Post("/search_availability", handlers.Repo.PostAvailability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Post("/search_availability_json", handlers.Repo.JsonTest)
	//setting up fileserver to handle static files (images etc.)
	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileserver))
	return mux
}
