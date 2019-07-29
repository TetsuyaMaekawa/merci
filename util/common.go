package util

import "strings"

// SplitString 文字列分割
func SplitString(textMessage string) []string {
	return strings.Split(textMessage, " ")
}
