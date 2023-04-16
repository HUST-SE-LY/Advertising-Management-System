package gin_ext

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindJSON(c *gin.Context, obj any) (err error) {
	err = c.BindJSON(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response(err, nil))
	}
	return
}

func ShouldBindJSON(c *gin.Context, obj any) (err error) {
	err = c.ShouldBindJSON(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response(err, nil))
	}
	return
}
