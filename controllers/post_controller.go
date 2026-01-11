package controllers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func HiWorld(e *core.RequestEvent) error {
	return e.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World! desde controller - New push 09"})
}
