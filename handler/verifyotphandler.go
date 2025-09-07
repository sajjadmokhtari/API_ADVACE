package handler

import (
	"Web/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PhoneAndOtp struct {
	Phone string `json:"phone"`
	Otp   string `json:"otp"`
}

func VerifyHandler(c *gin.Context) {
	var req PhoneAndOtp

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, Response{
			Valid:   false,
			Message: "invalid input",
		})
		return
	}

	storedOtp, err := service.GetOtp(req.Phone)
	if err != nil {
		c.JSON(401, Response{
			Valid:   false,
			Message: "OTP منقضی شده یا وجود نداره",
		})
		return
	}
	fmt.Println("otp saved reach and save")

	if storedOtp != req.Otp {
		c.JSON(401, Response{
			Valid:   false,
			Message: "your otp is wrong",
		})
		return
	}
	fmt.Println("otp is valid your login")

	c.JSON(200, Response{
		Valid:   true,
		Message: "you loggin successfully",
	})
}
