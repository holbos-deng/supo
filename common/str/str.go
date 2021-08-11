package str

func MethodNameHump(str string) string {
	v := []rune(str)
	if str == "" {
		return ""
	}
	v[0] = upper(v[0])
	for i := 1; i < len(str); i++ {
		if str[i] == '_' {
			v = append(v[0:i], v[i+1:len(str)]...)
			v[i] = upper(v[i])
		}
	}
	return string(v)
}

func upper(b rune) rune {
	if b >= 97 && b <= 122 {
		b -= 32
	}
	return b
}
