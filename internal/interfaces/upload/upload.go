package upload

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Невозможно прочитать файл"})
		return
	}

	filename := header.Filename
	log.Println("Файл:", filename)

	path := "internal/static/f/" + c.PostForm("id") + "/"
	err = os.MkdirAll(path, 0755)
	if err != nil {
		c.JSON(500, gin.H{"error": "Невозможно создать папку"})
		return
	}

	err = c.SaveUploadedFile(header, path+filename)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, "/room/"+c.PostForm("id"))
}

func Close(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		panic("AAA")
	}

	err := os.RemoveAll("internal/static/f/" + id)
	if err != nil {
		panic(err)
	}
	c.Redirect(302, "/")
}
