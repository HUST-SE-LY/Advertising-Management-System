package gin_ext

import (
	"backend/utils/status"
	"github.com/gin-gonic/gin"
)

func Response(err error, data interface{}) gin.H {
	var st status.Status
	var ok bool
	if st, ok = err.(status.Status); ok {
		st = status.ErrorToStatus(err)
	}
	return gin.H{
		"code":    st.Code,
		"message": st.Msg,
		"data":    data,
	}
}
