package router

import (
	"ddb-go-demo/internal/controller"
	"ddb-go-demo/internal/server"
	"net/http"
)

func ServeRoutes(ctx server.Context) http.Handler {
	router := http.NewServeMux()
	c := controller.NewBaseController(ctx)

	router.HandleFunc("POST /soap", c.SoapController)
	router.HandleFunc("GET /hello", c.HelloWorld)
	router.HandleFunc("GET /getUser", c.GetUserByEmail)
	router.HandleFunc("POST /saveUser", c.SaveUser)

	return router
}
