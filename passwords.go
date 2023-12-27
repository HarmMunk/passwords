package passwords

type GeneratorParams struct {
	NGramFreq           map[string]int
	NGramCnt            map[string]int
	MaxNGramLen         int
	SpecialCharacterMap []byte
	Randomiser          func(max int) int
}

func maxNGramLen(m map[string]int) int {
	ml := 0
	for k := range m {
		ml = max(ml, len(k))
	}
	return ml
}

func (gp GeneratorParams) PickNGram(nGram string) (next byte) {
	if gp.NGramCnt[nGram] == 0 {
		return gp.PickNGram(nGram[1:])
	}
	rnd := gp.Randomiser(gp.NGramCnt[nGram])
	next = 'a'
	var accu = gp.NGramFreq[nGram+string(next)]
	for ; next <= 'z'; next++ {
		accu += gp.NGramFreq[nGram+string(next)]
		if accu > rnd {
			break
		}
	}
	return next
}

func (gp GeneratorParams) Generate(length int) (password string) {
	if gp.MaxNGramLen == 0 {
		gp.MaxNGramLen = maxNGramLen(gp.NGramFreq)
	}
	if gp.MaxNGramLen == 0 {
		return ""
	}
	password = ""
	for len(password) < length {
		password += string(gp.PickNGram(password[max(len(password)-gp.MaxNGramLen, 0):len(password)]))
	}
	return password
}
