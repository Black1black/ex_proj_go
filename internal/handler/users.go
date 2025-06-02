package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) myInfo(c *gin.Context) {
	rawUserId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found in context")
		return
	}

	userId, ok := rawUserId.(int64)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid user id type in context")
		return
	}

	user, err := h.usrUC.GetByID(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

type GetUserInfoResponse struct {
	ID        int     `json:"id"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

type UserRequest struct {
	// ID - идентификатор пользователя (обязательный параметр пути)
	// Пример: /user123 (где 123 - ID)
	// Валидация:
	//   - required - параметр обязателен
	//   - numeric - должен быть числовым значением
	// Теги:
	//   - form:"-" - игнорировать в query-параметрах
	//   - uri:"id" - брать значение из параметра пути {id}
	ID        int     `form:"-" uri:"id" binding:"required,numeric"`
	Latitude  *string `form:"latitude" binding:"omitempty"`
	Longitude *string `form:"longitude" binding:"omitempty"`
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	var req UserRequest

	// Автоматический парсинг параметров пути и query
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// Ваша бизнес-логика
	response := gin.H{
		"id":        req.ID,
		"latitude":  req.Latitude,  // nil будет опущен в JSON
		"longitude": req.Longitude, // nil будет опущен в JSON
	}

	c.JSON(http.StatusOK, response)
}

// 	users.PATCH("/users")
// 	users.PATCH("/location")
