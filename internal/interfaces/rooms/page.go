package rooms

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoomPage(c *gin.Context) {
	id := c.Param("id")
	list_of_files, _ := GetFilesInDirectory(id)
	c.HTML(http.StatusOK, "room.html", gin.H{
		"Id":    id,
		"Files": list_of_files,
	})
}

func GetFilesInDirectory(dir string) ([]string, error) {
	files, err := ioutil.ReadDir("internal/static/f/" + dir)
	if err != nil {
		return []string{}, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}
