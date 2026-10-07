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

// GetEvents godoc
// @Summary      Get list of events
// @Description  Get list of events with optional filter parameters
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        filter  query     dto.EventFilterParam  false  "Filter parameters"
// @Success      200     {object}  dto.Response
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /events [get]
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

// GetEventByIdHandler godoc
// @Summary      Get event detail by ID
// @Description  Get detailed information of a specific event by its ID
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /events/{id} [get]
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

// JoinEventHandler godoc
// @Summary      Join or unjoin an event
// @Description  Join or leave an event by event ID (requires authentication)
// @Tags         Events
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /events/{id} [post]
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

// GetUpcomingEventHandler godoc
// @Summary      Get upcoming events
// @Description  Retrieve a list of upcoming events
// @Tags         Events
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /events/upcoming [get]
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

// GetMyEventsHandler godoc
// @Summary      Get user's joined events
// @Description  Get list of events joined by the authenticated user
// @Tags         Events
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @Success      200  {object}  dto.Response
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /events/myevents [get]
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
