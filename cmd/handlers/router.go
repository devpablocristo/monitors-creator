package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
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
