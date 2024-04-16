package tests

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"testing"
	"users_example/cmd/devapi/furyapp"
	"users_example/internal/platform/environment"
)

var framework *gin.Engine

func TestMain(m *testing.M) {
	dependencies, _ := furyapp.BuildDependencies(environment.Development)
	fakeApp := furyapp.Build(dependencies)

	framework = fakeApp
	m.Run()
}

func TestPing(t *testing.T) {
	res := performRequest("GET", "/ping", "", framework, nil)
	assert.Equal(t, http.StatusOK, res.Code)
}
