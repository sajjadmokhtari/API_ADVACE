package router

import (
	"Web/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouter()*gin.Engine {
	r := gin.Default()


	//r.POST("/login",handler.LoginHandler)
	r.POST("/description",handler.AcceptDec)
	r.GET("/descriptions", handler.GetAllDescriptions)

	r.POST("/send-otp",handler.SendOtpHandler)
	r.POST("/verify-otp",handler.VerifyHandler)



	

	return r
}
