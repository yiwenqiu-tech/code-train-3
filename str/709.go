package str

import "strings"

func toLowerCase(s string) string {
	//bs := []byte(s)
	//for i := range bs {
	//	if bs[i] >= 'A' && bs[i] <= 'Z' {
	//		bs[i] = bs[i] + 'a' - 'A'
	//	}
	//}
	//return string(bs)
	return strings.ToLower(s)
}
