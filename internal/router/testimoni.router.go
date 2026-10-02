package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initTestimoniRouter(ctx *gin.Engine, db *pgxpool.Pool) {
	testiRouter := ctx.Group("/testimony")

	tr := repo.NewTestimoniRepo(db)
	nr := repo.NewNotifRepo(db)
	ts := service.NewTestimoniService(tr, nr)
	th := handler.NewTestimoniHandler(ts)

	testiRouter.GET("", th.GetAllTestimony)
	testiRouter.POST("/create", middleware.CheckToken, th.CreateTestimony)
}
