package handler

import (
    "Web/service"
    "fmt"
    "github.com/gin-gonic/gin"
)

type PhoneRequest struct {
    Phone string `json:"phone"`
}

type Response struct {
    Valid   bool   `json:"valid"`
    Message string `json:"message"`
}

func SendOtpHandler(c *gin.Context) {
    var req PhoneRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, Response{
            Valid:   false,
            Message: "شماره ارسال نشده یا نامعتبر است",
        })
        return
    }
    fmt.Println("mobile is get")

    if !service.CheckPhone(req.Phone) {
        c.JSON(400, Response{
            Valid:   false,
            Message: "شماره موبایل نامعتبر است",
        })
        return
    }
    fmt.Println("mobile is valid")

    otp := service.GenerateOtp()


    
    fmt.Printf("📱 شماره: %s | 🔐 OTP: %s\n", req.Phone, otp)

   
    err := service.SaveOtp(req.Phone, otp)
    if err != nil {
        c.JSON(500, Response{
            Valid:   false,
            Message: "خطا در ذخیره OTP",
        })
        return
    }
    fmt.Println("otp save in redis")
 
    c.JSON(200, Response{
        Valid:   true,
        Message: "OTP ساخته شد و ارسال شد (برای تست: " + otp + ")",
    })
    fmt.Println("otp is send")
}
