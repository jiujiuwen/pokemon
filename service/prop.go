// 宝可梦属性
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/models"
	"strconv"
)

// add
func CreateProp(c *gin.Context) {
	Prop := models.Prop{Name: c.PostForm("name")}
	if s, err := models.AddProp(&Prop); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "create Prop success", "Id": s.ID})
	}
}

// get all Props ,total 19 ,no pagination
func FetchAllProp(c *gin.Context)  {
	s := models.FetchAllProp()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}

// fetch one Prop
func FetchOneProp(c *gin.Context)  {
	PropId, _ := strconv.Atoi(c.Param("id"))
	s := models.FetchOneProp(PropId)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Prop found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}

// update
func UpdateProp(c *gin.Context) {
	propId, _ := strconv.Atoi(c.Param("id"))
	propName := c.PostForm("name")

	//PropRestraint, _ := strconv.Atoi(c.PostForm("restraint"))

	propRestraint := c.PostFormArray("restraint_id")
	//propRestraintInt := make([]int, 3)
	var propRestraintInt []int

	for _, v := range propRestraint {
		restraintId, _ := strconv.Atoi(v)
		propRestraintInt = append(propRestraintInt, restraintId)
	}

	prop := models.Prop{Name: propName, RestraintId: propRestraintInt}

	s := models.UpdateProp(propId, &prop)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Prop found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Prop updated successfully!", "data": s, "test": propRestraint})
}

// delete
func DeleteProp(c *gin.Context)  {
	PropId, _ := strconv.Atoi(c.Param("id"))
	isExist := models.DeleteProp(PropId)


	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Prop found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Prop deleted successfully!"})
}
