package main

import (
	"os"
	"time"

	"github.com/pquerna/otp/totp"

	"fmt"
)

func main() {
	time := time.Now()
	secret := os.Getenv("OTP_SECRET")
	validateOpts := totp.ValidateOpts{
		Period:    300,
		Skew:      0,
		Digits:    6,
		Algorithm: 0,
	}
	argsWithProg := os.Args
	if secret != "" && len(argsWithProg) < 2 {
		fmt.Println("No OTP Secret input from STDIN and OTP_SECRET not set")
		os.Exit(1)
	} else if len(argsWithProg) > 1 {
		secret = os.Args[1]
	}
	test, err := totp.GenerateCodeCustom(secret, time, validateOpts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(test)
}
