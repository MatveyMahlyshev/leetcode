package functions

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	newString := ""
	var negative bool
	if x >= 0 {
		negative = false
	} else {
		negative = true
	}
	if negative {
		x = x * -1
	}
	for x != 0 {
		newString += strconv.Itoa(x % 10)
		x = x / 10
	}

	x, _ = strconv.Atoi(newString)
	if negative {
		x = x * -1
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return x
}

func MyAtoi(s string) int {
	newString := ""
	wasDigit := false
	wasZero := false
	for _, v := range s {
		// если буква, то выходим
		if v > 64 && v < 91 || v > 96 && v < 123 {
			break
		}
		// если точка
		if v == 46 {
			break
		}
		// если знак после цифры, то выходим
		if (v == 45 || v == 43) && wasDigit {
			break
		}
		if (v == 45 || v == 43) && wasZero {
			return 0
		}
		// если попадается пробел, а новая строка непустая, то выходим
		if v == 32 && newString != "" {
			break
		}
		// если попалась цифра, то
		if v > 48 && v < 58 {
			// если попался ноль и цифр до этого ещё не было, то пропускаем его
			// иначе добавляем цифру к новой строке и помечаем, что встретили цифру
			newString += string(v)
			wasDigit = true
			continue
		}
		if v == 48 {
			wasZero = true
			if wasDigit {
				newString += string(v)
			}
			continue
		}

		// если знак, то добавляем его
		if v == 45 || v == 43 {
			newString += string(v)
		}

		if v < 48 && v > 57 {
			break
		}

	}
	newInt, _ := strconv.Atoi(newString)
	if newInt > math.MaxInt32 {
		return math.MaxInt32
	}
	if newInt < math.MinInt32 {
		return math.MinInt32
	}

	return newInt
}

func PrefixCount(words []string, pref string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) < len(pref) {
			continue
		}
		innerCount := 0
		for j := 0; j < len(pref); j++ {
			if words[i][j] == pref[j] {
				innerCount += 1
			}
		}
		if innerCount == len(pref) {
			count += 1
		}
	}
	return count
}
