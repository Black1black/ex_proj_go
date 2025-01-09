package handlers

import (
	"ex_proj_go/internal/db"
	"ex_proj_go/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")
	{
		product.GET("", GetProducts)
		product.POST("", CreateProduct)
		product.PUT("/:id", UpdateProduct)
		product.PATCH("/:id", PartialUpdateProduct)
		product.DELETE("/:id", DeleteProduct)
	}
}

// @Summary Get all products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	go func() {
		var products []models.Product
		db.DB.Find(&products)
		c.JSON(http.StatusOK, products)
	}()
}

// @Summary Create a product
// @Produce json
// @Param product body models.Product true "Product"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	go func() {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Create(&product)
		c.JSON(http.StatusCreated, product)
	}()
}

// @Summary Update a product
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	go func() {
		id := c.Param("id")
		var product models.Product
		if err := db.DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Save(&product)
		c.JSON(http.StatusOK, product)
	}()
}

// @Summary Partial update a product
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Product
// @Router /products/{id} [patch]
func PartialUpdateProduct(c *gin.Context) {
	go func() {
		id := c.Param("id")
		var product models.Product
		if err := db.DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.DB.Save(&product)
		c.JSON(http.StatusOK, product)
	}()
}

// @Summary Delete a product
// @Produce json
// @Param id path int true "Product ID"
// @Success 204
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	go func() {
		id := c.Param("id")
		if err := db.DB.Delete(&models.Product{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	}()
}
