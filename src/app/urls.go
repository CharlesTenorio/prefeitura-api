package app

import (
	"apilogin/src/controllers"
	"prefeitura-api/src/controllers"
)

func mapUrls() {
	rota.GET("/ping", controllers.Ping)
}
