package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initUserRouter(r *gin.Engine, db *pgxpool.Pool) {
	userRouter := r.Group("/user")

	ur := repo.NewUserRepo(db)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)

	userRouter.GET("/my-profile", middleware.CheckToken, uh.GetProfileUserHandler)
	userRouter.PATCH("/change-profile", middleware.CheckToken, uh.ChangeUserProfile)
}
