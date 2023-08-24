package server

import (
	"github.com/gin-gonic/gin"

	"elm/internal/handler"
	"elm/internal/middleware"
	"elm/pkg/helper/resp"
	"elm/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler,
	homeHandler handler.HomeHandler,
	articleHandler handler.ArticleHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(log),
	)

	// No route group has permission
	noAuthRouter := r.Group("/api")
	{
		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
		noAuthRouter.POST("/check_login", userHandler.CheckLogin)
	}

	// Article rounter
	articleRouter := r.Group("/api")
	{
		articleRouter.GET("/articles/:id", articleHandler.GetArticleById)
		articleRouter.GET("/articles", articleHandler.GetArticleList)
	}
	articleAuthRouter := r.Group("/api").Use(middleware.StrictAuth(jwt, logger))
	{
		articleAuthRouter.POST("/articles", articleHandler.AddArticle)
		articleAuthRouter.POST("/article_content", articleHandler.AddArticleContent)
		articleAuthRouter.POST("/image/upload", articleHandler.ImageUpload)
		articleAuthRouter.POST("/article_content/update", articleHandler.UpdateArticleContent)
	}

	// home_page rounter
	homePageRouter := r.Group("/api")
	{
		homePageRouter.GET("/home_page", homeHandler.GetHomePage)
	}

	// Non-strict permission routing group
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
