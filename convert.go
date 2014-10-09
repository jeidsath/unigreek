package unigreek

import (
	"code.google.com/p/go.text/unicode/norm"
	"unicode/utf8"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

var greekUnicode = map[rune]rune{
	'a': 'α',
	'b': 'β',
	'g': 'γ',
	'd': 'δ',
	'e': 'ε',
	'z': 'ζ',
	'h': 'η',
	'q': 'θ',
	'i': 'ι',
	'k': 'κ',
	'l': 'λ',
	'm': 'μ',
	'n': 'ν',
	'c': 'ξ',
	'o': 'ο',
	'p': 'π',
	'r': 'ρ',
	's': 'σ',
	't': 'τ',
	'u': 'υ',
	'f': 'φ',
	'x': 'χ',
	'y': 'ψ',
	'w': 'ω',
        ':': '·',
        '\'': '’',
}

var markUnicodeInt = map[rune]int{
	'\\': 768, // Grave ὶ
	'/':  769, // Acute ί
	'+':  776, // Diaeresis ϊ
	')':  787, // Smooth breathing ἰ
	'(':  788, // Rough breathing ἱ
	'=':  834, // Circumflex ῖ
	'|':  837, // Iota subscript ᾳ
	//'':  , // Breve ῐ
	//'':  , // Macron ῑ
}

func unicodeIToS(ii int) string {
	/* TODO Understand the buffer here */
	bs := make([]byte, 2)
	_ = utf8.EncodeRune(bs, rune(ii))
	return string(bs)
}

func uppercase(input string) string {
	result := ""
	other := ""
	uppercase := false
	for _, runeValue := range input {
		switch {
		case runeValue == '*':
			uppercase = true
			continue
		case uppercase && runeValue >= 945 && runeValue <= 969:
			result += string(runeValue-32) + other
			uppercase = false
			other = ""
		case uppercase:
			other += string(runeValue)
		case !uppercase:
			result += string(runeValue)
		}
	}
	return result
}

func sigma(input string) string {
	result := ""
	sigma := false
	for _, runeValue := range input {
		if sigma {
			if (runeValue >= 945 && runeValue <= 969) || runeValue == '’' {
				result += "σ"
			} else {
				result += "ς"
			}
			sigma = false
		} 
                if runeValue == 'σ' {
                        sigma = true
                } else {
                        result += string(runeValue)
                }
	}
	if sigma {
		result += "ς"
	}
	return result
}

func Convert(input string) (string, error) {
	output := make([]string, len(input))

        // Ignore html escape
        offUntil := ""

        for ii, runeValue := range input {
                if offUntil != "" {
                        if offUntil == string(runeValue) {
                                offUntil = ""
                        }
                        output[ii] = string(runeValue)
                        continue
                }
                if string(runeValue) == "&" {
                        offUntil = ";"
                        output[ii] = string(runeValue)
                        continue
                }

                bytes := []byte(string(runeValue))
                if len(bytes) > 1 {
                        output[ii] = string(runeValue)
                        continue
                }
		if val, ok := greekUnicode[runeValue]; ok {
			output[ii] = string(val)
		} else {
			if val, ok := markUnicodeInt[runeValue]; ok {
				output[ii] += unicodeIToS(val)
			} else {
				output[ii] = string(runeValue)
			}
		}
	}

	result := ""
	for ii := 0; ii < len(output); ii++ {
		result += output[ii]
	}

	result = uppercase(result)
	result = sigma(result)

	return norm.NFC.String(result), nil
}
