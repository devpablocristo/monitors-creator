package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users_example/cmd/devapi/furyapp/handlers/presenter"
	"users_example/internal/supervisor"
)

type reportHandler struct {
	useCase supervisor.ReporterUseCase
}

func NewReportHandler(useCase supervisor.ReporterUseCase) *reportHandler {
	return &reportHandler{useCase: useCase}
}

func (h reportHandler) HandleLazyReport(c *gin.Context) {
	report, err := h.useCase.LazyReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusOK, presenter.Report(report))
}
