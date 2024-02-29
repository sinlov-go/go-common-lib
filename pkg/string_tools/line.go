package string_tools

import "strings"

func String2LineRaw(target string) string {
	newStr := strings.Replace(target, "\r\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\r", `\n`, -1)
	return newStr
}
