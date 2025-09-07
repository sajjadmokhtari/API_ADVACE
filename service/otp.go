package service

import (
	"Web/data/db"
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func CheckPhone(phone string) bool {

	re := regexp.MustCompile(`^09\d{9}$`)
	return re.MatchString(phone)

}

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000)) 
}

func SaveOtp(phone, otp string) error {//  سیو در ردیس 
	ctx := context.Background()
	key := "otp:" + phone
	ttl := 2 * time.Minute 

	return db.RedisClient.Set(ctx, key, otp, ttl).Err()
}
func GetOtp(phone string) (string, error) {
    ctx := context.Background()
    key := "otp:" + phone
    return db.RedisClient.Get(ctx, key).Result()
}