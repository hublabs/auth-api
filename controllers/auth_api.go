package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type AuthApiController struct {
}

func (c AuthApiController) Init(g *echo.Group) {
	g.GET("/api/v1/auths/ping", c.ping)
}

func (c AuthApiController) ping(ctx echo.Context) error {

	return ReturnResultApiSucc(ctx, http.StatusOK, "auth-ping")

}
