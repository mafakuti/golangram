package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "success")
	}
}
