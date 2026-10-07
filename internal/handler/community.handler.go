package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
)

type CommunityHandler struct {
	cs *service.CommunitiesService
}

func NewCommunityService(cs *service.CommunitiesService) *CommunityHandler {
	return &CommunityHandler{
		cs: cs,
	}
}

// GetCommunities godoc
// @Summary      Get list of communities
// @Description  Get list of communities with optional query filter parameters
// @Tags         Communities
// @Accept       json
// @Produce      json
// @Param        filter  query     dto.CommunityFilterParam  false  "Filter parameters"
// @Success      200     {object}  dto.Response
// @Failure      400     {object}  dto.ErrorResponse
// @Failure      401     {object}  dto.ErrorResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /community [get]
func (ch *CommunityHandler) GetCommunities(ctx *gin.Context) {
	var filter dto.CommunityFilterParam

	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	// val, exists := ctx.Get("token")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
	// 		Success: false,
	// 		Msg:     "unauthorized",
	// 	})
	// 	return
	// }

	// claims, ok := val.(pkg.JwtClaims)
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
	// 		Success: false,
	// 		Msg:     "terjadi kesalahan sistem",
	// 	})
	// 	return
	// }

	communities, err := ch.cs.GetCommunityService(ctx.Request.Context(), filter)
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

// GetCommunityByID godoc
// @Summary      Get community detail by ID
// @Description  Get detailed information of a specific community by its ID
// @Tags         Communities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Community ID"
// @Success      200  {object}  dto.Response
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /community/{id} [get]
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

	// val, exists := ctx.Get("token")
	// if !exists {
	// 	ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
	// 		Success: false,
	// 		Msg:     "unauthorized",
	// 	})
	// 	return
	// }

	// claims, ok := val.(pkg.JwtClaims)
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
	// 		Success: false,
	// 		Msg:     "terjadi kesalahan sistem",
	// 	})
	// }

	detail, err := ch.cs.GetCommunityDetailService(ctx.Request.Context(), communityID)
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

// GetPopularCommunities godoc
// @Summary      Get list of popular communities
// @Description  Retrieve a list of popular communities
// @Tags         Communities
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /community/popular [get]
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
