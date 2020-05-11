package album

import (
	"api/common/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbum(c *gin.Context) {
	albumIDString := c.Param("album_id")
	albumID, err := strconv.Atoi(albumIDString)
	if err != nil {
		panic(err)
	} else {
		c.JSON(200, dto.NewOKAPIResponse(FindAlbumByID(albumID)))
	}
}

func Register(g *gin.RouterGroup) {

	g.GET("/album/:album_id", getAlbum)
}
