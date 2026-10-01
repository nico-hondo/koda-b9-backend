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

func initNotifRouter(ctx *gin.Engine, db *pgxpool.Pool, redisClient *redis.Client) {
	notifRouter := ctx.Group("/my-notification")

	nr := repo.GetNewNotifRepo(db)
	ns := service.NewNotifService(nr, redisClient)
	nh := handler.NewNotifHandler(ns)

	notifRouter.GET("", middleware.CheckToken, nh.GetAllNotif)
}
