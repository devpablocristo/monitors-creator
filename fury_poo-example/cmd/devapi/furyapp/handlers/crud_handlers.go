package handlers

import (
	"net/http"
	"users_example/cmd/devapi/furyapp/handlers/presenter"
	"users_example/internal/supervisor"

	"github.com/gin-gonic/gin"
)

type crudHandler struct {
	useCase supervisor.CRUDUseCase
}

func NewCRUDHandler(useCase supervisor.CRUDUseCase) *crudHandler {
	return &crudHandler{useCase: useCase}
}

type createDeveloperRequest struct {
	Name      string `json:"name" binding:"required"`
	Team      string `json:"team" binding:"required"`
	Seniority string `json:"seniority" binding:"required"`
}

func (h crudHandler) HandleCreate(c *gin.Context) {
	var payload createDeveloperRequest
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, presenter.ApiError{})
		return
	}

	dev, err := h.useCase.Create(payload.Name, payload.Team, payload.Seniority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusCreated, presenter.Developer(dev))
}

func (h crudHandler) HandleRead(c *gin.Context) {
	id := c.Param("id")
	dev, err := h.useCase.Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.ApiError{})
		return
	}

	c.JSON(http.StatusOK, presenter.Developer(dev))
}

func (h crudHandler) HandleUpdate(c *gin.Context) {

}

func (h crudHandler) HandleDelete(c *gin.Context) {

}
