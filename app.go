package main

import (
	codechal "codechal/generate"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No argument provided!")
		os.Exit(1)
	}

	input := os.Args[1]

	if codechal.IsValidInputLength(input) != true {
		os.Exit(1)
	}

	pubKey := codechal.ReadFile(fmt.Sprintf("%s_public.pem", input))
	signature := codechal.ReadFile(fmt.Sprintf("%s_signature.txt", input))

	if pubKey == "" && signature == "" {
		signature, pubKey = codechal.GenerateNewKeys(input)
	}

	codechal.PrettyPrintJSON(input, signature, pubKey)

}
