package emailutil

import (
	"math/rand"
	"time"
)

func GenerateCode() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		code[i] = byte('0' + rand.Intn(10))
	}
	return string(code)
}
