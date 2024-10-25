package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/tsawler/celeritas"
)

type application struct {
	App *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
