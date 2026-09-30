package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/pkg"
)

func CheckToken(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "please login first",
		})
		c.Abort()
		return
	}
	result := strings.Split(bearer, " ")
	if len(result) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "invalid bearer token",
		})
		c.Abort()
		return
	}
	if result[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "invalid bearer token",
		})
		c.Abort()
		return
	}
	var token pkg.JwtClaims
	err := token.DecodeToken(result[1])
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Msg:     "invalid token",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}
	c.Set("token", token)
	c.Next()
}
