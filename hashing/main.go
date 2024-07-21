package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var hash string = base64Encode("mugdho")
	var decoded string = base64Decode(hash)

	if base64Compare("mugdho", hash) {
		fmt.Println("matched")
	} else {
		fmt.Println("didn't matched")
	}

	fmt.Println("hashed string : " + hash)

	fmt.Println("decoded string : ", decoded)
}

func base64Encode(plain string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(plain))
}

func base64Decode(hash string) string {
	decoded, _ := base64.RawStdEncoding.DecodeString(hash)
	return string(decoded)
}

func base64Compare(plain string, hashed string) bool {
	return base64Decode(hashed) == plain
}
