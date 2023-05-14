package router

import (
	"final_project/controllers"
	"final_project/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	photoRouter := r.Group("photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.GET("/:ID", controllers.GetPhotoById)
		photoRouter.PUT("/:ID", middlewares.Authorization("Photo"), controllers.UpdatePhoto)
		photoRouter.DELETE("/:ID", middlewares.Authorization("Photo"), controllers.DeletePhoto)
		photoRouter.GET("/:ID/comments", middlewares.Authorization("Photo"), controllers.GetComment)
	}

	socialMediaRouter := r.Group("social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.GET("/:ID", controllers.GetSocialMediaById)
		socialMediaRouter.PUT("/:ID", middlewares.Authorization("SocialMedia"), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:ID", middlewares.Authorization("SocialMedia"), controllers.DeleteSocialMedia)
	}

	commentRouter := r.Group("comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/:photoID", controllers.CreateComment)
		// commentRouter.GET("/", controllers.GetComment)
		commentRouter.GET("/:ID", controllers.GetCommentById)
		commentRouter.PUT("/:ID", middlewares.Authorization("Comment"), controllers.UpdateComment)
		commentRouter.DELETE("/:ID", middlewares.Authorization("Comment"), controllers.DeleteComment)
	}

	return r
}
