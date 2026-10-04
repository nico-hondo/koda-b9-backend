package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type CommunityHandler struct {
	cs *service.CommunitiesService
}

func NewCommunityService(cs *service.CommunitiesService) *CommunityHandler {
	return &CommunityHandler{
		cs: cs,
	}
}

func (ch *CommunityHandler) GetCommunities(ctx *gin.Context) {
	var filter dto.CommunityFilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
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

	communities, err := ch.cs.GetCommunityService(ctx.Request.Context(), filter, claims.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    communities,
		Msg:     "Communities berhasil diambil",
	})
}

func (ch *CommunityHandler) GetCommunityByID(ctx *gin.Context) {
	communityIDParam := ctx.Param("id")
	communityID, err := strconv.Atoi(communityIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     "ID komunitas tidak valid",
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
	}

	detail, err := ch.cs.GetCommunityDetailService(ctx.Request.Context(), communityID, claims.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    detail,
		Msg:     "Detail komunitas berhasil diambil",
	})
}

func (ch *CommunityHandler) GetPopularCommunities(ctx *gin.Context) {
	communities, err := ch.cs.GetPopularCommunitiesService(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "Gagal mengambil data popular communities",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    communities,
		Msg:     "Berhasil mengambil data popular communities",
	})
}
