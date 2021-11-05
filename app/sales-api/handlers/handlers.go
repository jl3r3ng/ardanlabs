package handlers

import (
	"gitlab.com/FireH24d/business/auth"
	"gitlab.com/FireH24d/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/FireH24d/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))

	return app
}
