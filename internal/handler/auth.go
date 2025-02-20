package handler

import (
	"ex_proj_go/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) login(c *gin.Context) {

	// rawUserModel, ok := c.Get(userModelCtx) // c.Get возвращает any
	// if !ok {
	// 	return
	// }
	// userModel, ok := rawUserModel.(*entity.User)
	// if !ok {
	// 	return
	// }

	var request loginRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.authUC.GetIdByEmail(request.Login, authTools.GeneratePasswordHash(request.Password))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	refreshToken, err := authTools.GenerateToken(id, "refresh")
	if err != nil {
		return
	}
	accessToken, err := authTools.GenerateToken(id, "access")
	if err != nil {
		return
	}
	err = h.authUC.Login(id, refreshToken)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, auth.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
