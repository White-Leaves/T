package Md5

import (
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
)

func Md5encrypt(str string) string {
	salt, encodedPWD := password.Encode(str, nil)
	check := password.Verify(str, salt, encodedPWD, nil)
	return fmt.Sprintf("%x", check)
}
