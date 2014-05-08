package main

import (
	"github.com/GeertJohan/btcqr"
	"log"
	"os"
)

func main() {
	file, err := os.Create("example-payment.png")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	btcqr.DefaultConfig.Scheme = "guldencoin"

	req := &btcqr.Request{
		Address: "Gf2S5QyHiMTcBD8anLuRmz7b4mQTQWeSLD",
		Amount:  1.234,
		Label:   "donate to GeertJohan",
		Message: "Or don't",
	}

	code, err := req.GenerateQR()
	if err != nil {
		log.Fatalf("error generating qr code: %v", err)
	}

	png := code.PNG()
	_, err = file.Write(png)
	if err != nil {
		log.Fatalf("error writing png to file: %v", err)
	}
}
