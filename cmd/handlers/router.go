package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/melisource/fury_go-platform/pkg/fury"
)

func NewHTTPServer(h MonitorHandler) error {
	// port := env.GetEnv("HANDLER_PORT")
	port := ":9000"
	r := gin.Default()

	basePath := "/api"
	router := r.Group(basePath)

	router.POST("monitor", h.CreateMonitor)
	router.GET("monitor/:id", h.GetMonitor)
	router.PATCH("monitor", h.UpdateMonitor)
	router.DELETE("monitor", h.DeleteMonitor)
	router.GET("monitor", h.GetAllMonitors)

	log.Println("Server listening on port", port)

	// return http.ListenAndServe(port, r)
	return r.Run(port)
}

func NewFuryApplication(handler *MonitorHandler) error {
	app, err := fury.NewWebApplication()
	fmt.Println(app)
	if err != nil {
		return err
	}
	app.Method(http.MethodPost, "/teste", func(w http.ResponseWriter, r *http.Request) error {
		return web.EncodeJSON(w, "MEU OVO", http.StatusOK)
	})
	app.Get("/hello", func(w http.ResponseWriter, r *http.Request) error {
		name := r.URL.Query().Get("name")
		if name == "" {
			return web.NewError(http.StatusBadRequest, "param name cannot be empty")
		}

		type response struct {
			Message string `json:"message"`
		}

		return web.EncodeJSON(w, response{"Hello " + name}, http.StatusOK)
	})

	return app.Run()
}
