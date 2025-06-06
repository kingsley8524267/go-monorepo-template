package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
)

func toPascalCase(input string) string {
	re := regexp.MustCompile(`[ _\-]+`)
	normalized := re.ReplaceAllString(input, " ")

	var withSpaces strings.Builder
	for i, r := range normalized {
		if i > 0 && unicode.IsUpper(r) && (unicode.IsLower(rune(normalized[i-1])) || unicode.IsDigit(rune(normalized[i-1]))) {
			withSpaces.WriteRune(' ')
		}
		withSpaces.WriteRune(r)
	}

	caser := cases.Title(language.English)
	words := strings.Fields(withSpaces.String())
	for i, w := range words {
		words[i] = caser.String(strings.ToLower(w))
	}

	return strings.Join(words, "")
}

func toSnakeCase(input string) string {
	re := regexp.MustCompile(`[ \-]+`)
	s := re.ReplaceAllString(input, "_")

	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(s[i-1])) || unicode.IsDigit(rune(s[i-1]))) {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	normalized := regexp.MustCompile(`_+`).ReplaceAllString(result.String(), "_")

	return strings.Trim(normalized, "_")
}
