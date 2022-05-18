package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

type TokenParsing interface {
	ParseToken(accessToken string) (int, error)
	GetUserCtx() string
}

func Authorization(tokenParserService TokenParsing) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)

		if header == "" {
			c.JSON(http.StatusUnauthorized, "empty auth header")
			c.Abort()
			return
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, "invalid auth header")
			c.Abort()
			return
		}

		if len(headerParts[1]) == 0 {
			c.JSON(http.StatusUnauthorized, "token is empty")
			c.Abort()
			return
		}

		userId, err := tokenParserService.ParseToken(headerParts[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set(tokenParserService.GetUserCtx(), userId)
	}
}
