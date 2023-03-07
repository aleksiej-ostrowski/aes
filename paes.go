package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"

	// "crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

const (
	info = `

#---------------------------------#
#                                 #
#  version 0.0.3                  #
#                                 #
#  Aleksiej Ostrowski, 2021-2023  #
#                                 #
#  https://aleksiej.com           #
#                                 #
#---------------------------------#

For encrypting a file:

... | paes password_for_encrypting e > file_name_for_result_encrypting


For decrypting a file:

... | paes password_for_decrypting d > file_name_for_result_decrypting

`
)

func generateKeyIV(password string) ([]byte, []byte) {
	key := make([]byte, 32)
	iv := make([]byte, 16)
	hash := sha256.Sum256([]byte(password))
	copy(key, hash[:32])
	copy(iv, hash[32:])
	return key, iv
}

func crypt_file(inp *os.File, out *os.File, password string, mode byte) (int, error) {

	key, iv := generateKeyIV(password)

	block, err := aes.NewCipher(key)
	if err != nil {
		return 2, err
	}

	var stream cipher.Stream
	if mode == 1 {
		stream = cipher.NewCFBEncrypter(block, iv)
	} else if mode == 2 {
		stream = cipher.NewCFBDecrypter(block, iv)
	} else {
		return 3, fmt.Errorf("Error mode = %d", mode)
	}

	reader := bufio.NewReader(inp)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf[:n], buf[:n])
			if _, err := writer.Write(buf[:n]); err != nil {
				return 4, err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return 5, err
		}
	}

	return 0, nil
}

func main() {

	var (
		code int = -1
	)

	defer func() {

		switch code {
		case 0:
			fmt.Println("ok")
		case 1:
			fmt.Println(info)
		case -1:
			fmt.Println("Unknown error")
		default:
			fmt.Printf("Error %d\n", code)
		}

		os.Exit(code)
	}()

	if len(os.Args) != 3 {
		code = 1
		return
	}

	// rand.Seed(time.Now().UnixNano())

	password := os.Args[1]

	inp := os.Stdin
	defer inp.Close()

	out := os.Stdout
	defer out.Close()

	var mode byte = 0

	switch what := os.Args[2]; what {
	case "e":
		mode = 1
	case "d":
		mode = 2
	}

	return_code, err := crypt_file(inp, out, password, mode)
	if err != nil {
		code = return_code
		return
	}

	code = 0
}
