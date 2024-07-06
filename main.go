package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to encrypt: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert to Unicode and multiply by 2
	encrypted := make([]int, len(input))
	for i, c := range input {
		encrypted[i] = int(c) * 2
	}

	// If greater than 126, subtract by 94 until no value is greater than 126
	for i, v := range encrypted {
		for v > 126 {
			v -= 94
		}
		encrypted[i] = v
	}

	// Add 1 if the final output is 32
	for i, v := range encrypted {
		if v == 32 {
			encrypted[i] = v + 1
		}
	}

	// Convert back to string
	var encryptedStr string
	for _, v := range encrypted {
		encryptedStr += string(rune(v))
	}

	// Perform MD5 hashing 4 times
	hash := md5Hash(encryptedStr)
	for i := 0; i < 3; i++ { // already hashed once, so 3 more times
		hash = md5Hash(hash)
	}

	fmt.Println("Encrypted text:", encryptedStr)
	fmt.Println("Final MD5 hash:", hash)
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
