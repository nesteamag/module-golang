package downcase

func Downcase(text string) (string, error) {
	buff := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		c := text[i]
		if (c >= 65) && (c <= 90) {
			c |= 'a' - 'A'
		}
		buff[i] = c
	}
	return string(buff), nil
}
