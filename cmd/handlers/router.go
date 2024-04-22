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

	log.Println("Server listening on port", port)

	// return http.ListenAndServe(port, r)
	return r.Run(port)
}
