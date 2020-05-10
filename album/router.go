package album

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbum(c *gin.Context) {
	albumIDString := c.Param("album_id")

	albumID, err := strconv.Atoi(albumIDString)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(200, FindAlbumByID(albumID))
	}

}

func Register(g *gin.RouterGroup) {
	g.GET("/album/:album_id", getAlbum)
}
