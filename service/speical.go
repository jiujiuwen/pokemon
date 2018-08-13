// 宝可梦特性
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/models"
	"strconv"
)

// add
func CreateSpecial(c *gin.Context) {
	Special := models.Special{Name: c.PostForm("name")}
	if s, err := models.AddSpecial(&Special); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "create Special success", "Id": s.ID})
	}
}

// get all Specials
// todo: pagination
func FetchAllSpecial(c *gin.Context)  {
	s := models.FetchAllSpecial()

	if s.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Special found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}
// fetch one Special
func FetchOneSpecial(c *gin.Context)  {
	SpecialId, _ := strconv.Atoi(c.Param("id"))
	s := models.FetchOneSpecial(SpecialId)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Special found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}

// update
func UpdateSpecial(c *gin.Context) {
	SpecialId, _ := strconv.Atoi(c.Param("id"))
	SpecialName := c.PostForm("name")

	s := models.UpdateSpecial(SpecialId, SpecialName)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Special found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Special updated successfully!", "data": s})
}

// delete
func DeleteSpecial(c *gin.Context)  {
	SpecialId, _ := strconv.Atoi(c.Param("id"))
	isExist := models.DeleteSpecial(SpecialId)


	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Special found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Special deleted successfully!"})
}
