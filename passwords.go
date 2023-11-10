package passwords

import (
	"regexp"
	"strconv"
)

var onlyConsonants, _ = regexp.Compile(`[aeiouy]{3}`)
var onlyVowels, _ = regexp.Compile(`[bcdfghjklmnpqrstvwxz]{3}`)

type GeneratorParams struct {
	NGramFreq           map[string]int
	SpecialCharacterMap []byte
	Randomiser          func(max int) int
	NGramCnt            map[string]int
}

func (gp GeneratorParams) PickNGram(nGram string) string {
	var rnd int
	if gp.NGramCnt[nGram] == 0 {
		return gp.PickNGram(nGram[1:])
	} else {
		rnd = gp.Randomiser(gp.NGramCnt[nGram])
	}
	var next byte = 'a'
	var accu int
	for next, accu = 'a', 0; next < 'z' && accu+gp.NGramFreq[nGram+string(next)] <= rnd; next, accu = next+1, accu+gp.NGramFreq[nGram+string(next)] {
	}
	return string(next)
}

func (gp GeneratorParams) Generate(length int) string {
	var password = ""

	password = gp.PickNGram("")
	password += gp.PickNGram(string(password[0]))
	for len(password) < length {
		i := len(password)
		ch := gp.PickNGram(password[i-2 : i])
		if onlyConsonants.MatchString(password[i-2:i]+ch) || onlyVowels.MatchString(password[i-2:i]+ch) {
			continue
		}
		password += ch
	}

	return password
}

func Unravel(ngram *map[string]int, s *string) {
	re, _ := regexp.Compile("([a-z]*)([0-9]*)")

	for _, v := range re.FindAllStringSubmatch(*s, -1) {
		(*ngram)[v[1]], _ = strconv.Atoi(v[2])
	}
}
