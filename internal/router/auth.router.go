package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initAuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	authRouter := r.Group("/auth")

	ar := repo.NewAuthRepo(db)
	nr := repo.NewNotifRepo(db)
	as := service.NewAuthService(ar, nr)
	ah := handler.NewAuthHandler(as)

	authRouter.POST("register", ah.Register)
	authRouter.POST("", ah.Login)
	authRouter.POST("/change-password", middleware.CheckToken, ah.ChangePassword)
	authRouter.PATCH("/change-profile", middleware.CheckToken, ah.ChangeProfileUser)
}
