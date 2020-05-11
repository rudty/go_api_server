package middleware

import (
	"api/common/dto"

	"github.com/gin-gonic/gin"
)

func recoverFunction(c *gin.Context) {
	err := recover()
	if err != nil {
		c.JSON(500, dto.NewErrorAPIResponse(err))
	}
}

// Recovery panic 발생 시 클라이언트에게 에러를 반환합니다
func Recovery(c *gin.Context) {
	defer recoverFunction(c)
	c.Next()
}
