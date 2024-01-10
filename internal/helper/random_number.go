package helper

import (
	"fmt"
	"math/rand"
)

func GenerateOtpNumber() string {
	randomNumber := rand.Intn(10000)
	return fmt.Sprintf("%04d", randomNumber)
}
