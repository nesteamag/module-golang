package letter

func Frequency(str string) map[rune]int {
	m := make(map[rune]int)

	for _, val := range str {
		m[val]++
	}

	return m
}

func parallel_frequency(str string, c chan map[rune]int) {
	c <- Frequency(str)
}

func ConcurrentFrequency(arr_str []string) map[rune]int {
	c := make(chan map[rune]int)
	m := make(map[rune]int)

	for i := range arr_str {
		go parallel_frequency(arr_str[i], c)
	}

	for range arr_str {
		for letter, count := range <-c {
			m[letter] += count
		}
	}

	return m
}
