package controllers

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func Version(e *core.RequestEvent) error {
	return e.JSON(http.StatusOK, map[string]string{
		"message":     "Testing deploy alive",
		"version":     "15-mar-2026",
		"last update": "Testing deploy alive",
	})
}
