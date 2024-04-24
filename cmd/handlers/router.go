package handlers

import (
	"github.com/melisource/fury_go-platform/pkg/fury"
)

func NewFuryApplication(h *MonitorHandler) error {
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}
	CreateRoutes(h, app)
	return app.Run()
}
