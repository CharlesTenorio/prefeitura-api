package app

import "github.com/gin-gonic/gin"

var (
	rota = gin.Default()
)

func IniciarApp() {
	mapUrls()
	rota.Run(":8080")

}
