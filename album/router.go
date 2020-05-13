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

func addAlbum(c *gin.Context) {
	albumIDString := c.PostForm("album_id")
	title := c.PostForm("title")
	albumID, err := strconv.Atoi(albumIDString)

	if err != nil {
		panic(err)
	}

	album := Album{
		AlbumID: albumID,
		Title:   title,
	}

	InsertAlbum(&album)
	c.JSON(200, dto.NewOKAPIResponse("OK"))
}

func modifyAlbum(c *gin.Context) {
	albumIDString := c.Param("album_id")
	title := c.PostForm("title")
	albumID, err := strconv.Atoi(albumIDString)

	if err != nil {
		panic(err)
	}

	album := FindAlbumByID(albumID)
	album.Title = title
	SaveAlbum(album)
}

// Register 라우터 등록
func Register(g *gin.RouterGroup) {
	g.GET("/album/:album_id", getAlbum)
	g.POST("/album", addAlbum)
	g.PUT("/album/:album_id", modifyAlbum)

}
