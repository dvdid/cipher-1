package main

import (
	"bufio"
	"fmt"
	"github.com/dvdid/cipher-1/encode"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(encode.Encode(scanner.Text()))
	}
}
