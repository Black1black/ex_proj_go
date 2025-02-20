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
	tokenCtx            = "token"
	userModelCtx        = "userModel"
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

	c.Set(tokenCtx, token)
}
func getUserId(c *gin.Context) {
	rawToken, ok := c.Get(tokenCtx)
	if !ok {
		return
	}
	token, ok := rawToken.(string)
	if !ok {
		return
	}
	userId, err := authTools.ParseToken(token)
	if err != nil {
		return
	}
	c.Set(userCtx, userId)

}
