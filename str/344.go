package str

func reverseString(s []byte) {
	bc := len(s)
	halfBc := bc / 2
	for i := 0; i < halfBc; i++ {
		s[i], s[bc-1-i] = s[bc-1-i], s[i]
	}
}
