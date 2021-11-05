package handlers

import (
	"gitlab.com/tleuzhan13/service/business/auth"
	"gitlab.com/tleuzhan13/service/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/tleuzhan13/service/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness, middleware.Authenticate(a), middleware.Authorize(auth.RoleAdmin))

	return app
}
