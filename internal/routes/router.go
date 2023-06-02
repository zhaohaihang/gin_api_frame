package routes

import (
	"net/http"

	v1 "gin_api_frame/internal/api/v1"

	"gin_api_frame/pkg/middleware"

	"gin_api_frame/internal/valid"

	_ "gin_api_frame/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter( uc *v1.UserController) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("../static"))

	valid.Init()

	v1 := r.Group("api/v1")
	{
		v1.GET("/swagger/*any", func(c *gin.Context) {
			ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
		})

		v1.POST("user/login", uc.UserLogin)
		v1.GET("user/:uid", uc.ViewUser)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", uc.UserUpdate)
			authed.PUT("user/changepasswd", uc.ChangePasswd)
			authed.POST("user/avatar", uc.UploadUserAvatar)
		}
	}
	return r
}

var RouterProviderSet = wire.NewSet(NewRouter)
