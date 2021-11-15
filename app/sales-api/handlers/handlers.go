package handlers

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/FireH24d/business/auth"
	"gitlab.com/FireH24d/business/data/book"
	"gitlab.com/FireH24d/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/FireH24d/foundation/web"
)

// API constructs an http.Handler with all application routes defined.

func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) http.Handler {
	// Construct the web.App which holds all routes as well as common Middleware.

	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))
	// Register debug check endpoints.

	cg := checkGroup{
		build: build,
		db:    db,
	}
	app.HandleDebug(http.MethodGet, "/readiness", cg.readiness)
	app.HandleDebug(http.MethodGet, "/liveness", cg.liveness)
	//hello
	bg := bookGroup{
		book: book.New(log, db),
	}
	// Register user management and authentication endpoints.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))

	http.HandleFunc("/", bg.handleListBooks)
	http.HandleFunc("/book.html", bg.handleViewBook)
	http.HandleFunc("/save", bg.handleSaveBook)
	http.HandleFunc("/delete", bg.handleDeleteBook)
	log.Fatal(http.ListenAndServe(":8080", nil))
	return app
}
