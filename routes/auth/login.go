package auth

import (
	"fmt"

	"github.com/gominima/minima"
)

func Login() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		auth := req.GetHeader("Authorization")

		fmt.Printf("Logged in user with token: %s\n", auth)

		res.JSON(map[string]interface{}{
			"ok":      true,
			"message": "logged in",
		})
	}
}
