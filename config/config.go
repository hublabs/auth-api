package config

import (
	"os"
)

var (
	AppEnv            string
	Httpport          string
	Service           string
	DataBaseDriver    string
	AuthApiConnection string
)

func Read() {
	if appEnv := os.Getenv("APP_ENV"); appEnv == "" {
		AppEnv = "test"
	} else {
		AppEnv = appEnv
	}
	if httpport := os.Getenv("HTTP_PORT"); httpport == "" {
		Httpport = "8003"
	} else {
		Httpport = httpport
	}
	if dataBaseDriver := os.Getenv("DATABASE_DRIVER"); dataBaseDriver == "" {
		DataBaseDriver = "sqlite3"
	} else {
		DataBaseDriver = dataBaseDriver
	}
	if authApiConnection := os.Getenv("AUTH_API_CONNECTION"); authApiConnection == "" {
		AuthApiConnection = ":memory:"
	} else {
		AuthApiConnection = authApiConnection
	}
	Service = "auth-api"
}

func ReadForTest() {
	AppEnv = "test"
	Httpport = "8003"
	Service = "auth-api"
	DataBaseDriver = "sqlite3"
	AuthApiConnection = ":memory:"
}
