package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type EventsHandler struct {
	es *service.EventsService
}

func NewEventsHandler(es *service.EventsService) *EventsHandler {
	return &EventsHandler{
		es: es,
	}
}

func (eh *EventsHandler) GetEvents(ctx *gin.Context) {
	var filter dto.EventFilterParam
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	events, err := eh.es.GetEventsService(ctx.Request.Context(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    events,
		Msg:     "Events berhasil diambil",
	})

}

func (eh *EventsHandler) GetEventByIdHandler(ctx *gin.Context) {
	evId := ctx.Param("id")

	id, err := strconv.Atoi(evId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     "id event tidak valid",
		})
		return
	}

	eventDetail, err := eh.es.GetEventDetailService(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     fmt.Sprintf("gagal mengambil detail event: %v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    eventDetail,
		Msg:     fmt.Sprintf("Detail Event dengan id %d, berhasil diambil", id),
	})
}

func (eh *EventsHandler) JoinEventHandler(ctx *gin.Context) {
	evIdParam := ctx.Param("id")

	evId, err := strconv.Atoi(evIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     "id event tidak valid",
		})
		return
	}

	val, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Msg:     "unauthorized",
		})
		return
	}

	claims, ok := val.(pkg.JwtClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	isJoined, msg, err := eh.es.JoinEventsService(ctx.Request.Context(), claims.Id, evId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"is_joined": isJoined,
		},
		Msg: msg,
	})
}

func (eh *EventsHandler) GetUpcomingEventHandler(ctx *gin.Context) {
	// Panggil layer Service
	events, err := eh.es.GetUpcomingEventService(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "Gagal mengambil data upcoming event",
		})
		return
	}

	// Return response OK
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    events,
		Msg:     "Berhasil mengambil data upcoming event",
	})
}

func (eh *EventsHandler) GetMyEventsHandler(ctx *gin.Context) {
	val, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Msg:     "unauthorized",
		})
		return
	}

	claims, ok := val.(pkg.JwtClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	myEvents, err := eh.es.GetMyEventsService(ctx.Request.Context(), claims.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    myEvents,
		Msg:     "MyEvent berhasil diambil",
	})
}
