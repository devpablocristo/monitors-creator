package handlers

import (
	"net/http"

	"github.com/melisource/fury_go-core/pkg/web"

	"monitors-creator/internal/monitors-creator/monitor"
)

type MonitorHandler struct {
	usecase monitor.MonitorUsecasePort
}

func NewMonitorHandler(usecase monitor.MonitorUsecasePort) *MonitorHandler {
	return &MonitorHandler{usecase: usecase}
}

func (h MonitorHandler) CreateMonitor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload MonitorPayload

	if err := web.DecodeJSON(r, &payload); err != nil {
		//TODO MYLOG
		return web.EncodeJSON(w, err, http.StatusUnprocessableEntity)
	}

	monitor, err := mapPayloadToDomain(payload)
	if err != nil {
		//TODO MYLOG
		return web.EncodeJSON(w, err, http.StatusUnprocessableEntity)
	}

	if err := h.usecase.CreateMonitor(ctx, &monitor); err != nil {
		//TODO MYLOG
		return web.EncodeJSON(w, err, http.StatusInternalServerError)
	}

	// TODO: CHAMAR DOMAIN2JSON OU PASSAR STRUCT DIRETO
	// TODO: MYLOG
	return web.EncodeJSON(w, monitor, http.StatusOK)
}

func (h MonitorHandler) GetMonitor(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id := web.Param(r, "id")

	monitor, err := h.usecase.GetMonitor(ctx, id)
	if err != nil {
		// TODO: MYLOG
		return web.EncodeJSON(w, err, http.StatusInternalServerError)
	}

	// TODO: CHAMAR DOMAIN2JSON OU PASSAR STRUCT DIRETO
	// TODO: MYLOG
	return web.EncodeJSON(w, monitor, http.StatusOK)
}

func (h MonitorHandler) UpdateMonitor(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h MonitorHandler) DeleteMonitor(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h MonitorHandler) GetAllMonitors(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	monitor, err := h.usecase.GetAllMonitors(ctx)
	if err != nil {
		// TODO: MYLOG
		return web.EncodeJSON(w, err, http.StatusInternalServerError)
	}

	// TODO: CHAMAR DOMAIN2JSON OU PASSAR STRUCT DIRETO
	// TODO: MYLOG
	return web.EncodeJSON(w, monitor, http.StatusOK)
}
