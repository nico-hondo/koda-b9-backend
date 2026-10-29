package handler

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

// GetProfileByIdHandler godoc
// @Summary      Get Profile detail by ID
// @Description  Get detailed information of a specific profile user by its ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Security	 BearerToken
// @Success      200  {object}  dto.Response
// @Failure		 401  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /user/my-profile [get]
func (uh *UserHandler) GetProfileUserHandler(ctx *gin.Context) {
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

	myProfile, err := uh.us.GetProfileUserService(ctx.Request.Context(), claims.Id)

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
		Data:    myProfile,
		Msg:     "Data Profile User berhasil diambil.",
	})
}

// ChangeProfileUser godoc
// @Summary      Change User Profile
// @Description  Update user profile details including name, location, bio, and avatar image
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerToken
// @Param        nama    formData  string  false  "User full name"
// @Param        loc     formData  string  false  "User location"
// @Param        bio     formData  string  false  "User biography"
// @Param        avatar  formData  file    false  "User profile picture"
// @Success      200     {object}  dto.Response
// @Failure      400     {object}  dto.Response "Invalid request form body"
// @Failure      401     {object}  dto.Response "Unauthorized"
// @Failure      500     {object}  dto.Response "Internal server error"
// @Router       /user/change-profile [patch]
func (uh *UserHandler) ChangeUserProfile(ctx *gin.Context) {
	var body dto.UpdateProfile

	if err := ctx.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		log.Println("bind error:", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "invalid request form body",
		})
		return
	}

	var imgUrl string

	// ini untuk nge-handle upload file gambar
	if body.Img != nil {
		filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), "user", path.Ext(body.Img.Filename))

		// Tetap simpan file ke direktori fisik backend
		savePath := path.Join("public", "img", filename)

		if err := ctx.SaveUploadedFile(body.Img, savePath); err != nil {
			log.Println("save file error:", err.Error())
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Msg:     "terjadi kesalahan saat menyimpan foto",
			})
			return
		}

		// KUNCI: Simpan path ke DB yang ramah URL static baru
		imgUrl = fmt.Sprintf("img/%s", filename)
	}

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

	if err := uh.us.ChangeUserProfile(ctx.Request.Context(), body, claims.Id, imgUrl); err != nil {
		log.Println("service error:", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Msg:     "Profile user berhasil diperbarui",
	})
}
