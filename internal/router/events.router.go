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

func initEventsRouter(ctx *gin.Engine, db *pgxpool.Pool, redisClient *redis.Client) {
	eventsRouter := ctx.Group("/events")

	er := repo.NewEventsRepo(db)
	nr := repo.NewNotifRepo(db)
	es := service.NewEventsService(er, nr, redisClient)
	eh := handler.NewEventsHandler(es)

	eventsRouter.GET("", eh.GetEvents)
	eventsRouter.GET("/:id", eh.GetEventByIdHandler)
	eventsRouter.GET("/upcoming", eh.GetUpcomingEventHandler)
	eventsRouter.POST("/:id", middleware.CheckToken, eh.JoinEventHandler)
	eventsRouter.GET("/myevents", middleware.CheckToken, eh.GetMyEventsHandler)
}
