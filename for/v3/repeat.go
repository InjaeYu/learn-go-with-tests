package iteration

import (
	"bytes"
	"strings"
)

func Repeat(character string) string {
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteString(character)
	}
	return buffer.String()
}

func Repeat2(character string) string {
	return strings.Repeat(character, 5)
}

func Repeat3(character string) string {
	result := make([]string, 5)
	for i := range result {
		result[i] = character
	}
	return strings.Join(result, "")
}
