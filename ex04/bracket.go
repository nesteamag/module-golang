package brackets

import "strings"

func Bracket(str string) (bool, error) {
	type1, type2, type3 := "()", "[]", "{}"

	flag := true

	for flag {
		if strings.Contains(str, type1) || strings.Contains(str, type2) || strings.Contains(str, type3) {
			str = strings.Replace(strings.Replace(strings.Replace(str, type1, "", -1), type2, "", -1), type3, "", -1)
		} else if str == "" {
			return flag, nil
		} else {
			flag = false
		}
	}
	return flag, nil
}
