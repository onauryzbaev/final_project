package comment

import "strings"

// CensoredWords - список нецензурных слов
var CensoredWords = []string{"badword1", "badword2"}

// IsCensored - проверяет, содержит ли комментарий нецензурные слова
func IsCensored(comment string) bool {
	for _, word := range CensoredWords {
		if strings.Contains(strings.ToLower(comment), strings.ToLower(word)) {
			return true
		}
	}
	return false
}
