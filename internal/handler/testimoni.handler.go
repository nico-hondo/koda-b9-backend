package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type TestimoniHandler struct {
	ts *service.TestimoniService
}

func NewTestimoniHandler(ts *service.TestimoniService) *TestimoniHandler {
	return &TestimoniHandler{
		ts: ts,
	}
}

func (th *TestimoniHandler) GetAllTestimony(ctx *gin.Context) {
	testimonies, err := th.ts.GetTestimoniService(ctx.Request.Context())
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    testimonies,
		Msg:     "Testimoni berhasil diambil",
	})
}

func (th *TestimoniHandler) CreateTestimony(ctx *gin.Context) {
	var body dto.NewTestimoni
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Msg:     "invalid request body",
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

	if err := th.ts.NewCreateTestimoniService(ctx.Request.Context(), claims.Id, body.Comment); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Msg:     "Testimoni berhasil dibuat",
	})
}
