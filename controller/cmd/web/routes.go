package web

import (
	"net/http"

	"github.com/Besufikad17/minab_events/cmd/web/handlers"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()
	handlers := handlers.NewHandler(app.Ctx)

	router.GET("/", handlers.Ping)

	// auth routes
	router.POST("/auth/login", handlers.Login)
	router.POST("/auth/signup", handlers.Signup)

	return router
}
