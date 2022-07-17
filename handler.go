package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
)

func root(c echo.Context) error {
	res, _ := exec.Command("promtool").Output()
	return c.String(http.StatusOK, string(res))
}

func promtool(c echo.Context) error {
	args := strings.Split(c.Param("*"), "/")
	res, _ := exec.Command("promtool", args...).Output()
	return c.String(http.StatusOK, string(res))
}
