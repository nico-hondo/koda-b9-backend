package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nico-hondo/internal/dto"
	apperror "github.com/nico-hondo/internal/error"
	"github.com/nico-hondo/internal/service"
	"github.com/nico-hondo/pkg"
)

type AuthHandler struct {
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{
		as: as,
	}
}

// RegistrationAccount godoc
// @Summary      Register an Account
// @Description  Regist Acc from body
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      dto.RegisterRequest  true  "body to register"
// @Success      201   {object}  dto.Response
// @Failure      500   {object}  dto.ErrorResponse
// @Router       /auth/register [post]
func (ah *AuthHandler) Register(ctx *gin.Context) {
	var body dto.RegisterRequest

	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())

		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan error",
		})
		return
	}

	// gunakan service
	if err := ah.as.CreateUser(ctx.Request.Context(), body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Data:    body,
		Msg:     "User Registered",
	})
}

// SignInAccount godoc
// @Summary		Login an Account
// @Description	Login Acc from body
// @Tags		Auth
// @Accept		json
// @Produce		json
// @Param		body	body		dto.LoginRequest	true	"body to login"
// @Success		200		{object}	dto.Response
// @Failure		400		{object}	dto.ErrorResponse
// @Failure		401		{object}	dto.ErrorResponse
// @Failure		500		{object}	dto.ErrorResponse
// @Router		/auth	[post]
func (a *AuthHandler) Login(c *gin.Context) {
	var body dto.LoginRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan server",
		})
		return
	}
	// gunakan service
	token, err := a.as.Login(c.Request.Context(), body)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, apperror.ErrCredentialsEmpty) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Msg:     err.Error(),
			})
			return
		}
		if errors.Is(err, apperror.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Msg:     err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"token": token,
		},
	})
}

func (h *AuthHandler) LogoutHandler(ctx *gin.Context) {
	// Ambil token & claims dari gin context (disimpan oleh AuthMiddleware)
	tokenString, _ := ctx.Get("token")
	claims, _ := ctx.Get("claims")

	err := h.as.Logout(ctx.Request.Context(), tokenString.(string), claims.(jwt.MapClaims))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to logout",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully logged out",
	})
}

// ChangePassword godoc
// @Summary      Change Password User
// @Description  Change password for logged-in user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BearerToken
// @Param        request  body      dto.ChangePasswordRequest  true  "Data Password Baru"
// @Success      200      {object}  dto.Response
// @Failure      400      {object}  dto.Response
// @Failure      401      {object}  dto.Response
// @Failure      500      {object}  dto.Response
// @Router       /auth/change-password [patch]
func (a *AuthHandler) ChangePassword(c *gin.Context) {
	var body dto.ChangePasswordRequest
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "invalid request body",
		})
		return
	}

	val, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "unauthorized",
		})
		return
	}

	claims, ok := val.(pkg.JwtClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	if err := a.as.ChangePassword(c.Request.Context(), claims.Id, body.OldPassword, body.NewPassword); err != nil {
		log.Println(err.Error())
		if errors.Is(err, apperror.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Msg:     "password lama tidak sesuai",
			})
			return
		}
		if errors.Is(err, apperror.ErrCredentialsEmpty) {
			c.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Msg:     err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Msg:     "Password berhasil diganti",
	})
}

// ChangeProfileUser godoc
// @Summary      Change Profile User
// @Description  Update user profile details using Form Data
// @Tags         Auth
// @Accept       mpfd
// @Produce      json
// @Security     BearerToken
// @Param        body  body      dto.ChangeProfileUser  true  "Profile Data"
// @Success      200   {object}  dto.Response
// @Failure      400   {object}  dto.Response
// @Failure      401   {object}  dto.Response
// @Failure      500   {object}  dto.Response
// @Router       /auth/change-profile [patch]
func (ah *AuthHandler) ChangeProfileUser(ctx *gin.Context) {
	var body dto.ChangeProfileUser

	if err := ctx.ShouldBindWith(&body, binding.FormPost); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "invalid request body",
		})
		return
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

	if err := ah.as.ChangeUserProfile(ctx.Request.Context(), claims.Id, body.Name, body.Avatar_url, body.Bio, body.Location, body.Job, body.Workplace); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data:    body,
		Msg:     "Profile User berhasil diganti",
	})
}
