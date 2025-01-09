package handlers

import (
	"ex_proj_go/internal/db"
	"ex_proj_go/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	{
		user.GET("", GetUsers)
		user.POST("", CreateUser)
		user.PUT("/:id", UpdateUser)
		user.PATCH("/:id", PartialUpdateUser)
		user.DELETE("/:id", DeleteUser)
	}
}

// @Summary Get all users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	go func() {
		var users []models.User
		db.DB.Find(&users)
		c.JSON(http.StatusOK, users)
	}()
}

// @Summary Create a user
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	go func() {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Create(&user)
		c.JSON(http.StatusCreated, user)
	}()
}

// @Summary Update a user
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	go func() {
		id := c.Param("id")
		var user models.User
		if err := db.DB.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Save(&user)
		c.JSON(http.StatusOK, user)
	}()
}

// @Summary Partial update a user
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [patch]
func PartialUpdateUser(c *gin.Context) {
	go func() {
		id := c.Param("id")
		var user models.User
		if err := db.DB.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Save(&user)
		c.JSON(http.StatusOK, user)
	}()
}

// @Summary Delete a user
// @Produce json
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	go func() {
		id := c.Param("id")
		if err := db.DB.Delete(&models.User{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	}()
}
