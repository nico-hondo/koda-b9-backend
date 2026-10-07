package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/middleware"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool, redisClient *redis.Client) {
	// Pasang middleware CORS di sini secara global
	router.Use(middleware.Cors)

	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	initAuthRouter(router, db, redisClient)
	initNotifRouter(router, db, redisClient)
	initTestimoniRouter(router, db)
	initEventsRouter(router, db, redisClient)
	initCommunityRouter(router, db)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Msg:     "Page not found!",
		})
	})
}
