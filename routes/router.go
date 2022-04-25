package routes

import (
	"github.com/gominima/minima"
	"github.com/twisttaan/SplashBackend/routes/auth"
)

func Router() *minima.Group {
	return minima.NewGroup("").
		Get("/auth/login", auth.Login())
}
