package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const URL = "https://prnt.sc/"

func letterAndNumberGenerator(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return strings.ToLower(string(b))
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	go func() {
		for i := 0; i < 256*256+256; i++ {
			fmt.Println(URL + letterAndNumberGenerator(6))
		}
	}()
	var input string
	fmt.Scanln(&input)
	main()
}
