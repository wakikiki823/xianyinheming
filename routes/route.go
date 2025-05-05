package routes

import (
	"wenxinhexuan/controllers"
	"wenxinhexuan/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	userController *controllers.UserController,
	songController *controllers.SongController,
) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("/send_code", controllers.SendVerificationCode)    // 发送验证码
		userRoutes.POST("/verify_code", userController.VerifyCode)        // 验证验证码
		userRoutes.GET("/forget_password", userController.ForgetPassword) // 发送忘记密码请求
		userRoutes.GET("/reset_password", userController.ResetPassword)
	}

	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	{
		api.GET("/song", songController.GetSong)                // 需要在路由上加入歌曲ID参数
		api.GET("/update_list_song", songController.UpdateList) // 更新歌曲列表
		api.GET("/all_songs", songController.GetAllSongs)
		api.GET("/essay", userController.GetGeneratedEssay)
	}

	return r
}
