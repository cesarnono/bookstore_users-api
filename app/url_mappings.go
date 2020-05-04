package app

import (
	"github.com/cesarnono/bookstore_users-api/controllers/ping"
	"github.com/cesarnono/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
}
