package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initCommunityRouter(ctx *gin.Engine, db *pgxpool.Pool) {
	communityRouter := ctx.Group("/community")

	cr := repo.NewCommunityRepo(db)
	nr := repo.NewNotifRepo(db)
	cs := service.NewCommunityService(cr, nr)
	ch := handler.NewCommunityService(cs)

	communityRouter.GET("", middleware.CheckToken, ch.GetCommunities)
	communityRouter.GET("/:id", middleware.CheckToken, ch.GetCommunityByID)
	communityRouter.GET("/popular", ch.GetPopularCommunities)
}
