package middleware

import (
	"myapp/data"

	"github.com/tsawler/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}