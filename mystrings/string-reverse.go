package mystrings

import (
	"strings"
	"bytes"
)

func ReverseString(s string) string {
	//Return s if given string is empty or blank
	if len(s)==0  || len(strings.TrimSpace(s)) == 0{
		return s
	}

	var buffer bytes.Buffer
	for i:=len(s)-1;i>=0;i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}
