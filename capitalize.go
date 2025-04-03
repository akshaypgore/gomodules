package capitalize

import "strings"

func GetCapitalized(inputWord string) string {
	return strings.ToUpper(inputWord[0:1]) + strings.ToLower(inputWord[1:])
}
