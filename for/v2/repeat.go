package iteration

import "bytes"

func Repeat(character string) string {
	var buffer bytes.Buffer
	for i := 0; i < 5; i++ {
		buffer.WriteString(character)
	}
	return buffer.String()
}
