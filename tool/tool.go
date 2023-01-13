// Package tool /**
package tool

import (
	"log"
	"strings"
)

// InitialToCapital 首字母转大写
func InitialToCapital(str string) string {
	log.Println(str)

	arr := strings.Split(str, "_")

	var (
		resultStr string
	)
	for _, s := range arr {
		var InitialToCapitalStr string
		strRune := []rune(s)
		for i := 0; i < len(strRune); i++ {
			if i == 0 {
				if strRune[i] >= 97 && strRune[i] <= 122 {
					strRune[i] -= 32
					InitialToCapitalStr += string(strRune[i])
				} else {
					return s
				}
			} else {
				InitialToCapitalStr += string(strRune[i])
			}
		}
		//return InitialToCapitalStr
		resultStr += InitialToCapitalStr
	}

	log.Println(resultStr)

	return resultStr
}

// PathProcessing 路径处理
func PathProcessing(path string) bool {
	a := strings.Split(path, "/")
	if len(a) >= 2 && a[1] != "" {
		return false
	} else {
		return true
	}

}
