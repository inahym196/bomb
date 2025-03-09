package shared

import (
	"fmt"
	"strings"
	"unicode"
)

func NumToExcelColumn(n int) string {
	result := ""
	for n > 0 {
		n--
		result = string(rune('A'+(n%26))) + result
		n /= 26
	}
	return result
}

func ExcelColumnToNum(s string) (int, error) {
	result := 0
	for _, ch := range strings.ToUpper(s) {
		if !unicode.IsLetter(ch) {
			return 0, fmt.Errorf("%c is not letter", ch)
		}
		result = result*26 + (int(ch - 'A'))
	}
	return result, nil
}
