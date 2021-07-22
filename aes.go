// https://stackoverflow.com/questions/52696921/reading-bytes-into-go-buffer-with-a-fixed-stride-size
// https://levelup.gitconnected.com/a-short-guide-to-encryption-using-go-da97c928259f

package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

const (
	info = `

#--------------------------------#
#                                #
#  version 0.0.2                 #
#                                #
#  Aleksiej Ostrowski, 2021      #
#                                #
#  https://aleksiej.com          #
#                                #
#--------------------------------#

For encrypting a file:

./aes file_name_for_encrypting password_for_encrypting e


For decrypting a file:

./aes file_name_for_decrypting password_for_decrypting d

`
)

var (
	code int = -1
)

func crypt_file(inp *os.File, out *os.File, password string, mode byte) {

	// encrypt = [mode = 1]
	// decrypt = [mode = 2]

	hash := sha256.Sum256([]byte(password))
	// fmt.Printf("%x", hash[:])

	block, err := aes.NewCipher(hash[:])
	if err != nil {
		code = 2
		panic(err)
	}

	iv := make([]byte, block.BlockSize())
	buf := make([]byte, 1024)

	if mode == 1 {
		rand.Read(iv)
	}

	var msgLen int64

	if mode == 2 {

		st, err := inp.Stat()
		if err != nil {
			code = 4
			panic(err)
		}

		msgLen = st.Size() - int64(len(iv))

		_, err = inp.ReadAt(iv, msgLen)
		if err != nil {
			code = 5
			panic(err)
		}
	}

	stream := cipher.NewCTR(block, iv)

	w := bufio.NewWriter(out)

	for {

		n, err := inp.Read(buf)

		flag := n > 0

		if flag {

			if mode == 2 {

				if n > int(msgLen) {
					n = int(msgLen)
				}

				msgLen -= int64(n)
			}

			stream.XORKeyStream(buf, buf[:n])
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			if err != io.ErrUnexpectedEOF {
				code = 6
				panic(err)
			}
		}

		if flag {
			_, err = w.Write(buf[:n])

			if err != nil {
				code = 7
				panic(err)
			}
		}
	}

	if mode == 1 {
		_, err = w.Write(iv)

		if err != nil {
			code = 13
			panic(err)
		}
	}

	err = w.Flush()
	if err != nil {
		code = 8
		panic(err)
	}

}

func main() {

	defer func() {

		switch code {
		case 0:
			fmt.Println("ok")
		case -1:
			fmt.Println("Unknown error")
		case 1:
			fmt.Println(info)
		default:
			fmt.Printf("Error %v\n", code)
		}

		if r := recover(); r != nil {
			fmt.Println("Recovering from panic: ", r)
		}

		os.Exit(code)
	}()

	if len(os.Args) != 4 {
		code = 1
		return
	}

	// rand.Seed(time.Now().UnixNano())

	input_filename := os.Args[1]

	password := os.Args[2]

	output_filename := input_filename + "_"

	inp, err := os.Open(input_filename)
	if err != nil {
		code = 9
		return
	}
	defer inp.Close()

	st, err := inp.Stat()
	if err != nil {
		code = 10
		return
	}

	if st.Size() == 0 {
		code = 11
		return
	}

	out, err := os.Create(output_filename)
	if err != nil {
		code = 12
		return
	}
	defer out.Close()

	var mode byte = 0

	switch what := os.Args[2]; what {
	case "e":
		mode = 1
	case "d":
		mode = 2
	}

	crypt_file(inp, out, password, mode)

	code = 0
}
