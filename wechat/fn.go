package wechat

func reverseStr(s string) string {
	r := []rune(s)
	l := len(r)
	nr := make([]rune, l)
	for i := 0; i < l; i++ {
		nr[l-i-1] = r[i]
	}
	return string(nr)
}
