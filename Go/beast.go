/**
 * Syc <github.com/SycAlright>
 * Beast_SDK GoLang
 */

package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

var beast = []string{"嗷", "呜", "啊", "~"}

func main() {
	enStr := encode("你好")
	fmt.Print(enStr)
	deStr := decode("呜嗷嗷嗷啊嗷嗷~啊呜~啊~呜呜嗷")
	fmt.Print(deStr)
}

func decode(str string) string {
	array := strings.Split(str, "")
	var code string
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			if array[i+1] == "" {
				break
			}
			a := search(array[i])
			b := search(array[i+1])
			x := ((a * 4) + b) - ((i / 2) % 16)
			if x < 0 {
				x += 16
			}

			code += toHex(x)
		}
	}
	hex, _ := hex.DecodeString(code)
	return string(hex)
}

func encode(str string) string {
	hex := bin2hex(str)
	array := strings.Split(hex, "")
	var code string
	for i := 0; i < len(array); i++ {
		x := hex2dec(array[i]) + (i % 16)
		if x >= 16 {
			x -= 16
		}
		code += beast[(x/4)] + beast[x%4]
	}
	return code
}

func bin2hex(str string) string {
	src := []byte(str)
	dst := hex.EncodeToString(src)
	var output string
	for i := 0; i < len(dst); i++ {
		if i%2 == 0 {
			binstr := dst[i : i+2]
			output += binstr
		}
	}
	fmt.Println(output)
	return output
}

func hex2dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

func search(str string) int {
	for i := 0; i < len(beast); i++ {
		if beast[i] == str {
			return i
		}
	}
	return -1
}

func toHex(ten int) string {
	m := 0
	hex := make([]int, 0)
	for {
		m = ten % 16
		ten = ten / 16
		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	hexStr := []string{}
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexStr = append(hexStr, fmt.Sprintf("%c", 'A'+hex[i]-10))
		} else {
			hexStr = append(hexStr, fmt.Sprintf("%d", hex[i]))
		}
	}
	return strings.Join(hexStr, "")
}
