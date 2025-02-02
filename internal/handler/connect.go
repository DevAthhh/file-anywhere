package handler

import (
	view "github.com/DevAthhh/fileanywhere/internal/interfaces/main"
	"github.com/gin-gonic/gin"
)

func Handle() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("internal/templates/*")
	router.Static("static/", "internal/static/")

	view.LoadAll(router)

	return router
}
