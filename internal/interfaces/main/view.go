package view

import (
	view "github.com/DevAthhh/fileanywhere/internal/interfaces/index"
	"github.com/DevAthhh/fileanywhere/internal/interfaces/rooms"
	"github.com/DevAthhh/fileanywhere/internal/interfaces/upload"
	"github.com/gin-gonic/gin"
)

func LoadAll(router *gin.Engine) {
	router.GET("/", view.Index)
	router.GET("/room/:id", rooms.RoomPage)
	router.POST("/save", upload.Upload)
	router.GET("/close", upload.Close)
}
