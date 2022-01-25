package main

import (
	"fmt"
	"test-module/helper"

	"github.com/badoux/checkmail"
)

func main() {
	helper.Write()
	fmt.Println("Main File.")

	var email = "gu.bat98@gmail.com"
	validMail := checkmail.ValidateFormat(email)

	fmt.Println(validMail)
}
