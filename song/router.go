package song

import (
	"api/common/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getSong(c *gin.Context) {
	songIDString := c.Param("song_id")
	songID, err := strconv.Atoi(songIDString)

	if err != nil {
		panic(err)
	}

	c.JSON(200, dto.NewOKAPIResponse(FindSongByID(songID)))
}

// Register 라우터 등록
func Register(g *gin.RouterGroup) {
	g.GET("/song/:song_id", getSong)
}
