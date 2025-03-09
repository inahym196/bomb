package shared

import "unicode"

func NumToExcelColumn(n int) string {
	result := ""
	for n > 0 {
		n--
		result = string(rune('A'+(n%26))) + result
		n /= 26
	}
	return result
}

func ExcelColumnToNum(s string) int {
	result := 0
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			result = result*26 + (int(ch-'A') + 1)
		}
	}
	return result
}
