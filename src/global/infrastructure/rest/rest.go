package rest

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lea55/BACKENDCANDITICKET/src/global/core"
	restconst "github.com/lea55/BACKENDCANDITICKET/src/global/infrastructure/rest/constants"
)

func AppStart() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	ac := core.GetAppConfig()
	restconst.SeUseCaseList()

	listenHandlers(r, ac.BaseUrl, ac)

	err := r.Run(":" + os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
