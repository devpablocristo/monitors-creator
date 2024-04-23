package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"monitors-creator/cmd/handlers/presenter"
	"monitors-creator/internal/monitors-creator/monitor"
)

type MonitorHandler struct {
	usecase monitor.MonitorUsecasePort
}

func NewMonitorHandler(usecase monitor.MonitorUsecasePort) *MonitorHandler {
	return &MonitorHandler{usecase: usecase}
}

func (h MonitorHandler) CreateMonitor(c *gin.Context) {
	ctx := c.Request.Context()

	var payload MonitorPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error(err))
		return
	}

	monitor, err := mapPayloadToDomain(payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, presenter.Error(err))
		return
	}

	if err := h.usecase.CreateMonitor(ctx, &monitor); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error(err))
		return
	}

	c.JSON(http.StatusCreated, MonitorResponse(monitor))
}

// func (h MonitorHandler) MonitorRead(c *gin.Context) {
// 	// id := c.Param("id")
// 	// dev, err := h.usecase.Read(id)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, presenter.ApiError{})
// 	// 	return
// 	// }

// 	// c.JSON(http.StatusOK, presenter.Developer(dev))
// }

// func (h MonitorHandler) MonitorUpdate(c *gin.Context) {

// }

// func (h MonitorHandler) MonitorDelete(c *gin.Context) {

// }
