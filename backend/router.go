package main

import (
	"github.com/zenazn/goji/web"

	"github.com/nazarnovak/jayway/backend/api"
)

func router() *web.Mux {
	mux := web.New()

	mux.Post("/api/robot", api.RobotHandler())

	return mux
}

