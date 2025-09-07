package handler



import (
    "Web/data/db"
    "Web/model"
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetAllDescriptions(c *gin.Context) {
    var descriptions []model.Description

	database := db.GetDb()

	if err :=  database.Find(&descriptions).Error; err != nil {
		c.JSON(500,gin.H{
			"message":"error to find",
		})
		return
	}


    c.JSON(http.StatusOK, gin.H{
        "data": descriptions,
    })
}
