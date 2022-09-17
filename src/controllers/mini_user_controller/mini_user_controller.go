package mini_user_controller

import (
	"apis/common/resp"

	"github.com/gin-gonic/gin"
)

func LoginWithWxCode(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		resp.RespParamErr(ctx)
		return
	}

}
