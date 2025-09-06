package reverse

import "bytes"

func reverse(s string) string {
	var b bytes.Buffer
	endIdx := len(s) - 1
	for i := endIdx; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}
