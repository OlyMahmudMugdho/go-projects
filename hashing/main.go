package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var hash string = base64Encode("mugdho")
	fmt.Println(hash)
}

func base64Encode(plain string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(plain))
}
