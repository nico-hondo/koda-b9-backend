package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type NotifHandler struct {
	ps *service.NotifService
}

func NewNotifHandler(ps *service.NotifService) *NotifHandler {
	return &NotifHandler{
		ps: ps,
	}
}

// GetAllNotif godoc
// @Summary Get All Notif Data
// @Description Get Notif Data from database
// @Tags notif
// @Produce json
// @Security BearerToken
// @Success 200 {object} dto.Response
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /my-notification [get]
func (nh *NotifHandler) GetAllNotif(ctx *gin.Context) {

	val, exists := ctx.Get("token")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "unauthorized",
		})
		return
	}

	claims, ok := val.(pkg.JwtClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	notif, err := nh.ps.GetAllNotifbyId(ctx.Request.Context(), claims.Id)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    notif,
		Msg:     "Data Notif berhasil diambil.",
	})
}
