package controllers

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"

	configutil "github.com/hublabs/auth-api/config"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/kafka"
)

var (
	appEnv           = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	ctx              context.Context
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
	xormEngine       *xorm.Engine
)

func init() {
	runtime.GOMAXPROCS(1)
	configutil.ReadForTest()
	var err error
	xormEngine, err = xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	echoApp = echo.New()
	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return echomiddleware.ContextDB(configutil.Service, xormEngine, kafka.Config{})(handlerFunc)(c)
	}
	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}

func SetContext(req *http.Request) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, xormEngine.NewSession())))

	return c, rec
}
func SetContextWithSession(req *http.Request, session *xorm.Session) (echo.Context, *httptest.ResponseRecorder) {

	rec := httptest.NewRecorder()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, session)))

	return c, rec
}
