package handler

import (
	"Web/data/db"
	"Web/model"

	"github.com/gin-gonic/gin"
)

type DecRequest struct {
	Description string `json:"description"`
}

func AcceptDec(c *gin.Context) {
	var req DecRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(401, gin.H{
			"message": "error",
		})
		return

	}
	desc := model.Description{
		Description: req.Description,
	}

	database := db.GetDb()

	if err := database.Create(&desc).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "خطا در ذخیره‌سازی توضیحات",
		})

	}

	    c.JSON(200, gin.H{
        "message": "توضیحات با موفقیت ثبت شد!",
    })

}
