package handler

import (
	"ex_proj_go/internal/entity"
	"ex_proj_go/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) login(c *gin.Context) {

	var request loginRequest

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.authUC.GetIDByEmail(request.Login, authTools.GeneratePasswordHash(request.Password))
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
	c.SetCookie(authorizationCookie, accessToken, 3600*12, "", "", false, true)
	c.SetCookie(authorizationRefreshCookie, refreshToken, 3600*12, "", "", false, true)
	// имя: , значение: , срок жизни: , путь: , домен: "", secure: , HttpOnly:

	c.JSON(http.StatusOK, auth.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	rawUserModel, ok := c.Get(userModelCtx) // c.Get возвращает any
	if !ok {
		return
	}
	userModel, ok := rawUserModel.(*entity.User)
	if !ok {
		return
	}
	rawToken, ok := c.Get(tokenCtx)
	if !ok {
		return
	}
	token, ok := rawToken.(string)
	if !ok {
		return
	}

	err := h.authUC.DeleteRefreshToken(userModel.ID, token)
	if err != nil {
		return
	}

	refreshToken, err := authTools.GenerateToken(userModel.ID, "refresh")
	if err != nil {
		return
	}

	accessToken, err := authTools.GenerateToken(userModel.ID, "access")
	if err != nil {
		return
	}
	err = h.authUC.Login(userModel.ID, refreshToken)
	if err != nil {
		return
	}
	c.SetCookie(authorizationCookie, accessToken, 3600*12, "", "", false, true)
	c.SetCookie(authorizationRefreshCookie, refreshToken, 3600*12, "", "", false, true)
	// имя: , значение: , срок жизни: , путь: , домен: "", secure: , HttpOnly:

	c.JSON(http.StatusOK, auth.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *Handler) logout(c *gin.Context) {

	rawToken, ok := c.Get(tokenCtx)
	if !ok {
		return
	}
	token, ok := rawToken.(string)
	if !ok {
		return
	}

	rawUserModel, ok := c.Get(userModelCtx) // c.Get возвращает any
	if !ok {
		return
	}
	userModel, ok := rawUserModel.(*entity.User)
	if !ok {
		return
	}

	err := h.authUC.DeleteRefreshToken(userModel.ID, token)
	if err != nil {
		return
	}
	c.SetCookie(authorizationCookie, "", -1, "", "", false, true)
	c.SetCookie(authorizationRefreshCookie, "", -1, "", "", false, true)

}

// type signInInput struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func (h *Handler) signIn(c *gin.Context) {
// 	var input signInInput

// 	if err := c.BindJSON(&input); err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 	}
// 	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
// }
