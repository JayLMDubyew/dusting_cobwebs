package main

import (
	"flag"
	"fmt"
)

func main() {
	var textFlag = flag.String("text", "", "Text to convert to cipherText")
	var shiftFlag = flag.Int("shift", 0, "how far to shift the letters in the string. default: 0")
	var decodeFlag = flag.Bool("decode", false, "attempt to decode")

	flag.Parse()

	fmt.Println(*textFlag, *shiftFlag)
	runes := []rune(*textFlag)
	for i, v := range runes {
		if !*decodeFlag {
			shifted := int(v) + *shiftFlag
			if (v >= 90 && v < 97) || (v > 122) || (v < 65) {
				continue
			}
			if (shifted >= 90 && shifted < 97) || (shifted > 122) {
				v -= 26
			}
			runes[i] = rune(int(v) + *shiftFlag)
		} else {

			shifted := int(v) - *shiftFlag
			if (v >= 90 && v < 97) || (v > 122) || (v < 65) {
				continue
			}
			if (shifted <= 64) || (shifted < 97 && shifted > 90) {
				v += 26
			}
			runes[i] = rune(int(v) - *shiftFlag)
		}
	}

	result := string(runes)
	fmt.Println(result)

	xor_mode()
}

func xor_mode() {
	key := byte('a')
	text := "helloz"
	newrunes := []byte(text)
	for i, v := range newrunes {
		encoded := key ^ v
		newrunes[i] = encoded

	}
	fmt.Printf("%q\n", string(newrunes))
	for i, v := range newrunes {
		decoded := key ^ v
		newrunes[i] = decoded

	}
	fmt.Printf("%q\n", string(newrunes))

}

