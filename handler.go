package main

import (
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func promtool(c echo.Context) error {
	res, err := exec.Command("promtool").Output()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(res))
}
