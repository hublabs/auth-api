package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/auth-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_AuthApiController_GetAuthById(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/v1/auths/ping", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, AuthApiController{}.ping(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  string                 `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result, "ping")
}
