package util

// Byte2Str ==> byte array to string
func Byte2Str(bytes ...byte) string {
	str := ""
	for _, v := range bytes {
		str += string(rune(v))
	}
	return str
}
