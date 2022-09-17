package routers

import (
	"apis/common/tool"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Run() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Content-Type,Authorization,RoleId,Uid,x-requested-with",
		ExposedHeaders:  "Content-Type,Authorization,RoleId,Uid,x-requested-with",
		Credentials:     false,
		ValidateHeaders: false,
	}))
	engine.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"msg": "ok"})
	})

	engine.GET("uuid", func(ctx *gin.Context) {
		_, _ = ctx.Writer.WriteString(fmt.Sprintf("%d", tool.GetSnowflakeId()))
	})

	engine.GET("time", func(context *gin.Context) {
		_, _ = context.Writer.WriteString(fmt.Sprintf("%d", time.Now().Unix()))
	})
	RegisteClient(engine)
	return engine
}
