package handlers

import (
	"gitlab.com/FireH24d/business/auth"
	"gitlab.com/FireH24d/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/FireH24d/foundation/web"
)

// API constructs an http.Handler with all application routes defined.

func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	// Construct the web.App which holds all routes as well as common Middleware.

	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))
	// Register debug check endpoints.

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))

	return app
}
