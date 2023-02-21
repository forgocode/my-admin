package v1

import "github.com/gin-gonic/gin"

func TestRespose(c *gin.Context) {
	c.JSON(200, "test successfully!")
}
