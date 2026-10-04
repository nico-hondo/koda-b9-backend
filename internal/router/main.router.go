package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitMainRouter(router *gin.Engine, db *pgxpool.Pool, redisClient *redis.Client) {

	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	initAuthRouter(router, db)
	initNotifRouter(router, db, redisClient)
	initTestimoniRouter(router, db)
	initEventsRouter(router, db)
	initCommunityRouter(router, db)
}
