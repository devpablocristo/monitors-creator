package handlers

import (
	"net/http"

	"github.com/melisource/fury_go-platform/pkg/fury"
)

type Routes struct {
	URL               string
	Method            string
	Function          func(w http.ResponseWriter, r *http.Request) error
	ReqAuthentication bool
}

func CreateRoutes(h *MonitorHandler, app *fury.Application) {
	var monitorRoutes = []Routes{
		{
			URL:               "/monitor",
			Method:            http.MethodPost,
			Function:          h.CreateMonitor,
			ReqAuthentication: false,
		},
		{
			URL:               "/monitor/{id}",
			Method:            http.MethodGet,
			Function:          h.GetMonitor,
			ReqAuthentication: false,
		},
		{
			URL:               "/monitor",
			Method:            http.MethodPatch,
			Function:          h.UpdateMonitor,
			ReqAuthentication: false,
		},
		{
			URL:               "/monitor",
			Method:            http.MethodDelete,
			Function:          h.DeleteMonitor,
			ReqAuthentication: false,
		},
		{
			URL:               "/monitor",
			Method:            http.MethodGet,
			Function:          h.GetAllMonitors,
			ReqAuthentication: false,
		},
	}

	for _, route := range monitorRoutes {

		// TODO if using middlewared:
		// if route.ReqAuthentication {
		// 	r.HandleFunc(route.URI,
		// 		middlewares.Logger(
		// 			middlewares.Autenticar(
		// 				route.Funcao))).Methods(route.Metodo)
		// 				} else {
		app.Method(route.Method, route.URL, route.Function)
	}
}
