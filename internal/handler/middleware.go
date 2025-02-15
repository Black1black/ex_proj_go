package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	authorizationCookie = "ttt"
	userCtx             = "userId"
)

func (h *Handler) getUserToken(c *gin.Context) {
	var token string

	cookie, err := c.Cookie(authorizationCookie)
	if err != nil && cookie != "" {
		token = cookie
	} else {
		header := c.GetHeader(authorizationHeader)

		headerParts := strings.Split(header, " ")
		token = headerParts[1]

		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}
	}

	c.Set(userCtx, token)
}
func getUserId(c *gin.Context) (int, error) {
	token, ok := c.Get(userCtx)

	return 0, nil
}

func getUserModel(c *gin.Context) (int, error) {
	token, ok := c.Get(userCtx)

	return 0, nil
}
