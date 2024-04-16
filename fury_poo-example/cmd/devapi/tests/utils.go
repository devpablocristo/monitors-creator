package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"net/http"
	"net/http/httptest"
	"strings"
)

func performRequest(method, target, body string, engine *gin.Engine, headers map[string]string) *httptest.ResponseRecorder {
	var req *http.Request
	if body == "" {
		req = httptest.NewRequest(method, target, http.NoBody)
	} else {
		payload := strings.NewReader(body)
		req = httptest.NewRequest(method, target, payload)
	}
	for headKey, headValue := range headers {
		req.Header.Add(headKey, headValue)
	}
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)
	logger.Infof(res.Body.String())
	return res
}
