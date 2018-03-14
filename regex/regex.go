package regex

import (
	"regexp"
)

func Get(content string, expr string) string {
	result := ""
	rule, _ := regexp.Compile(expr)
	match_results := rule.FindStringSubmatch(content)
	if 2 == len(match_results) {
		result = match_results[1]
	}
	return result
}

func GetAll(content string, expr string) (results []string) {
	rule, _ := regexp.Compile(expr)
	results = rule.FindAllString(content,-1)
	return results
}

func Replace(content string, expr string, replacement string) string {
	rule, _ := regexp.Compile(expr)
	return rule.ReplaceAllString(content, replacement)
}
