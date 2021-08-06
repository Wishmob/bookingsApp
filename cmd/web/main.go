package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Wishmob/bookingsApp/internal/config"
	"github.com/Wishmob/bookingsApp/internal/driver"
	"github.com/Wishmob/bookingsApp/internal/handlers"
	"github.com/Wishmob/bookingsApp/internal/helpers"
	"github.com/Wishmob/bookingsApp/internal/models"
	"github.com/Wishmob/bookingsApp/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	listenForMail()
	defer close(app.MailChan)

	// testMsg := models.MailData{
	// 	From:     "a@b.de",
	// 	To:       "c@d.com",
	// 	Subject:  "TestSubject",
	// 	Content:  "Hello <strong>WORLD</strong>",
	// 	Template: "basic.html",
	// }
	// app.MailChan <- testMsg

	defer db.SQL.Close()

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// change this to true when in production
	app.InProduction = false

	gob.Register(models.Reservation{}) //necessary to store data other than primitives in session
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})

	//set up error & infolog
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infolog
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//setting up mail chan
	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookingsApp user=alexbreiter password=")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
