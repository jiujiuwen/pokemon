// 宝可梦属性
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/models"
	"strconv"
)

// add
func CreateType(c *gin.Context) {
	Type := models.Type{Name: c.PostForm("name")}
	if s, err := models.AddType(&Type); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "create Type success", "Id": s.ID})
	}
}

// get all Types
// todo: pagination
func FetchAllType(c *gin.Context)  {
	s := models.FetchAllType()

	if s.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}
// fetch one Type
func FetchOneType(c *gin.Context)  {
	TypeId, _ := strconv.Atoi(c.Param("id"))
	s := models.FetchOneType(TypeId)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}

// update
func UpdateType(c *gin.Context) {
	TypeId, _ := strconv.Atoi(c.Param("id"))
	TypeName := c.PostForm("name")

	s := models.UpdateType(TypeId, TypeName)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Type updated successfully!", "data": s})
}

// delete
func DeleteType(c *gin.Context)  {
	TypeId, _ := strconv.Atoi(c.Param("id"))
	isExist := models.DeleteType(TypeId)


	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Type deleted successfully!"})
}
