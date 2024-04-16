package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users_example/cmd/devapi/furyapp/handlers/presenter"
	"users_example/internal/supervisor"
)

type tasksManagementHandler struct {
	useCase supervisor.TaskManagementUseCase
}

func NewTasksManagementHandler(useCase supervisor.TaskManagementUseCase) *tasksManagementHandler {
	return &tasksManagementHandler{useCase: useCase}
}

type scheduleTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h tasksManagementHandler) HandleScheduleTask(c *gin.Context) {
	var payload scheduleTaskRequest
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{})
		return
	}

	devId := c.Param("id")
	dev, err := h.useCase.ScheduleTask(devId, payload.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusOK, presenter.Developer(dev))
}

func (h tasksManagementHandler) HandleCompleteTask(c *gin.Context) {
	devId := c.Param("id")
	dev, err := h.useCase.CompleteTask(devId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusOK, presenter.Developer(dev))
}
