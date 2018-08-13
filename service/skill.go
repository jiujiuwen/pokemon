// 宝可梦技能
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/models"
	"strconv"
)

// add
func CreateSkill(c *gin.Context) {
	var propId, _ = strconv.Atoi(c.PostForm("prop_id"))
	Skill := models.Skill{Name: c.PostForm("name"), PropId: propId}

	if s, err := models.AddSkill(&Skill); err == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "create skill success", "Id": s.ID})
	}
}

// get all skills
func FetchAllSkill(c *gin.Context)  {
	pageNum, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("size"))
	name := c.Query("name")

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	s := models.FetchAllSkill(pageNum, pageSize, name)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": s,
		"total": models.GetTotalNum(),
		"page": pageNum,
		"size": pageSize,
		"name": name})
}
// fetch one skill
func FetchOneSkill(c *gin.Context)  {
	skillId, _ := strconv.Atoi(c.Param("id"))
	s := models.FetchOneSkill(skillId)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No skill found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": s})
}

// update
func UpdateSkill(c *gin.Context) {
	skillId, _ := strconv.Atoi(c.Param("id"))
	skillProp, _ := strconv.Atoi(c.PostForm("prop_id"))
	skill := models.Skill{Name: c.PostForm("name"), PropId: skillProp}

	s := models.UpdateSkill(skillId, &skill)

	if s.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Skill found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Skill updated successfully!", "data": s})
}

// delete
func DeleteSkill(c *gin.Context)  {
	skillId, _ := strconv.Atoi(c.Param("id"))
	isExist := models.DeleteSkill(skillId)


	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No skill found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Skill deleted successfully!"})
}
