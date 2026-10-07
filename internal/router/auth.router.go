package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
	"github.com/redis/go-redis/v9"
)

func initAuthRouter(r *gin.Engine, db *pgxpool.Pool, redisClient *redis.Client) {
	authRouter := r.Group("/auth")

	ar := repo.NewAuthRepo(db)
	nr := repo.NewNotifRepo(db)
	as := service.NewAuthService(ar, nr, redisClient)
	ah := handler.NewAuthHandler(as)

	authRouter.POST("register", ah.Register)
	authRouter.POST("", ah.Login)
	authRouter.PATCH("/change-password", middleware.CheckToken, ah.ChangePassword)
	authRouter.PATCH("/change-profile", middleware.CheckToken, ah.ChangeProfileUser)

	protect := r.Group("/")
	protect.Use(middleware.AuthMiddleware(redisClient, "YOUR_JWT_SECRET"))
	{
		protect.POST("/logout", ah.LogoutHandler)
	}
}
