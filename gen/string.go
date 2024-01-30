package gen

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"unicode"
)

var goKeyword = map[string]string{
	"var":         "variable",
	"const":       "constant",
	"package":     "pkg",
	"func":        "function",
	"return":      "rtn",
	"defer":       "dfr",
	"go":          "goo",
	"select":      "slt",
	"struct":      "structure",
	"interface":   "itf",
	"chan":        "channel",
	"type":        "tp",
	"map":         "mp",
	"range":       "rg",
	"break":       "brk",
	"case":        "caz",
	"continue":    "ctn",
	"for":         "fr",
	"fallthrough": "fth",
	"else":        "es",
	"if":          "ef",
	"switch":      "swt",
	"goto":        "gt",
	"default":     "dft",
}

// ToCamel 将字符串改为驼峰式
func ToCamel(s string) string {
	return ""
}

// Title 将字符串首字母改为大写
func Title(s string) string {
	if len(s) == 0 {
		return s
	}
	return cases.Title(language.English, cases.NoLower).String(s)
}

// Untitle 将字符串首字母改为小写
func Untitle(s string) string {
	if len(s) == 0 {
		return s
	}
	r := rune(s[0])
	if !unicode.IsUpper(r) && !unicode.IsLower(r) {
		return s
	}
	return string(unicode.ToLower(r)) + s[1:]
}

func wrapWithRawString(s string) string {
	return s
}

func EscapeGolangKeyword(s string) string {
	if !isGolangKeyword(s) {
		return s
	}

	r := goKeyword[s]
	return r
}

func isGolangKeyword(s string) bool {
	_, ok := goKeyword[s]
	return ok
}
