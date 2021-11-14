package handlers

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/FireH24d/business/auth"
	"gitlab.com/FireH24d/business/data/product"
	"gitlab.com/FireH24d/business/data/user"
	"gitlab.com/FireH24d/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/FireH24d/foundation/web"
)

// API constructs an http.Handler with all application routes defined.

func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {
	// Construct the web.App which holds all routes as well as common Middleware.

	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))
	// Register debug check endpoints.

	cg := checkGroup{
		build: build,
		db:    db,
	}

	app.Handle(http.MethodGet, "/readiness", cg.readiness)
	app.Handle(http.MethodGet, "/liveness", cg.liveness)
	ug := userGroup{
		user: user.New(log, db),
		auth: a,
	}
	app.Handle(http.MethodGet, "/users/:page/:rows", ug.query, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodGet, "/users/token/:kid", ug.token)
	app.Handle(http.MethodGet, "/users/:id", ug.queryByID, middleware.Authenticate(a))
	app.Handle(http.MethodPost, "/users", ug.create, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodPut, "/users/:id", ug.update, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))
	app.Handle(http.MethodDelete, "/users/:id", ug.delete, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))

	// Register product and sale endpoints.
	pg := productGroup{
		product: product.New(log, db),
	}
	app.Handle(http.MethodGet, "/v1/products/:page/:rows", pg.query, middleware.Authenticate(a))
	app.Handle(http.MethodGet, "/v1/products/:id", pg.queryByID, middleware.Authenticate(a))
	app.Handle(http.MethodPost, "/v1/products", pg.create, middleware.Authenticate(a))
	app.Handle(http.MethodPut, "/v1/products/:id", pg.update, middleware.Authenticate(a))
	app.Handle(http.MethodDelete, "/v1/products/:id", pg.delete, middleware.Authenticate(a))

	return app
}
