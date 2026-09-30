package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/handler"
	"github.com/nico-hondo/internal/middleware"
	"github.com/nico-hondo/internal/repo"
	"github.com/nico-hondo/internal/service"
)

func initAuthRouter(r *gin.Engine, db *pgxpool.Pool) {
	authRouter := r.Group("/auth")

	ar := repo.NewAuthRepo(db)
	nr := repo.NewNotifRepo(db)
	as := service.NewAuthService(ar, nr)
	ah := handler.NewAuthHandler(as)

	// authRouter.POST("pwd", func(ctx *gin.Context) {
	// 	type body struct {
	// 		Password string `json:"password"`
	// 	}

	// 	var reqBody body

	// 	if err := ctx.ShouldBindWith(&reqBody, binding.JSON); err != nil {
	// 		log.Println("error ", err.Error())
	// 		//binding
	// 		ctx.JSON(http.StatusInternalServerError, dto.Response{
	// 			Success: false,
	// 			Data:    nil,
	// 			Msg:     "terjadi kesalahan error",
	// 		})
	// 		return
	// 	}

	// 	hc := pkg.NewRecommendedHashConfig()
	// 	hash := hc.GenHash(reqBody.Password)

	// 	ctx.JSON(http.StatusOK, dto.Response{
	// 		Success: true,
	// 		Data: gin.H{
	// 			"pwd":  reqBody.Password,
	// 			"hash": hash,
	// 		},
	// 	})
	// })

	// authRouter.POST("compare", func(ctx *gin.Context) {
	// 	type body struct {
	// 		Password string `json:"pwd"`
	// 		Hash     string `json:"hash"`
	// 	}
	// 	var reqBody body
	// 	if err := ctx.ShouldBindWith(&reqBody, binding.JSON); err != nil {
	// 		log.Println("error", err.Error())
	// 		// binding error
	// 		ctx.JSON(http.StatusInternalServerError, dto.Response{
	// 			Success: false,
	// 			Data:    nil,
	// 			Msg:     "terjadi kesalahan server",
	// 		})
	// 		return
	// 	}

	// 	err := pkg.Compare(reqBody.Password, reqBody.Hash)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		if errors.Is(err, apperror.ErrInvalidHash) {
	// 			ctx.JSON(http.StatusUnauthorized, dto.Response{
	// 				Success: false,
	// 				Msg:     "password salah",
	// 			})
	// 			return
	// 		}
	// 		ctx.JSON(http.StatusInternalServerError, dto.Response{
	// 			Success: false,
	// 			Data:    nil,
	// 			Msg:     "terjadi kesalahan server",
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, dto.Response{
	// 		Success: true,
	// 		Msg:     "password betul",
	// 	})
	// })

	authRouter.POST("register", ah.Register)
	authRouter.POST("", ah.Login)
	authRouter.POST("/change-password", middleware.CheckToken, ah.ChangePassword)
	authRouter.PATCH("/change-profile", middleware.CheckToken, ah.ChangeProfileUser)
}
