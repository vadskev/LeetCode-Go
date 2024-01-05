package leetcode

var romanLettersMap = map[string]int{
	"M":  1000,
	"D":  500,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func RomanToInt(s string) int {
	numInt := 0
	for k := range s {
		if k < len(s)-1 && romanLettersMap[s[k:k+1]] < romanLettersMap[s[k+1:k+2]] {
			numInt -= romanLettersMap[s[k:k+1]]
		} else {
			numInt += romanLettersMap[s[k:k+1]]
		}
	}
	return numInt
}
