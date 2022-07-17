package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	fmt.Println("original secret", string(sd))
	whatIsIt = reverse(string(sd))
	fmt.Println("result with reverse", whatIsIt)
}

func reverse(s string) string {
	result := []rune(s)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
