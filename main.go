package main

import (
	"os"
	"time"

	"github.com/pquerna/otp/totp"

	"fmt"
)

func main() {
	time := time.Now()
	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		fmt.Println("No OTP Secret input")
		os.Exit(1)
	}
	secret := os.Args[1]
	test, err := totp.GenerateCode(secret, time)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(test)
}
