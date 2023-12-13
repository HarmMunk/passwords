package passwords

//var onlyConsonants, _ = regexp.Compile(`[aeiouy]{3}`)
//var onlyVowels, _ = regexp.Compile(`[bcdfghjklmnpqrstvwxz]{3}`)

type GeneratorParams struct {
	NGramFreq           map[string]int
	NGramCnt            map[string]int
	maxNGramLen         int
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
		//fmt.Println("Why did this happen?", nGram)
		return gp.PickNGram(nGram[1:])
	}
	//} else {
	//	rnd = gp.Randomiser(gp.NGramCnt[nGram])
	//}
	//fmt.Println("PickNGram 0: param nGram=", nGram)
	rnd := gp.Randomiser(gp.NGramCnt[nGram])
	//fmt.Println("PickNGram 1: rnd=", rnd, ", NGramCnt[", nGram, "]=", gp.NGramCnt[nGram])
	next = 'a'
	var accu = gp.NGramFreq[nGram+string(next)]
	//fmt.Println("PickNGram 2: accu=", accu)
	for ; next <= 'z'; next++ {
		accu += gp.NGramFreq[nGram+string(next)]
		if accu > rnd {
			break
		}
	}
	//fmt.Println("PickNGram 3: next=", string(next))
	//isViablePrefix := false
	//count := 0
	//for !isViablePrefix {
	//	if gp.NGramCnt[nGram] == 0 {
	//		return gp.PickNGram(nGram[1:len(nGram)])
	//	}
	//	rnd := gp.Randomiser(gp.NGramCnt[nGram])
	//	fmt.Println("PickNGram 1: rnd=", rnd, ", NGramCnt[", nGram, "]=", gp.NGramCnt[nGram])
	//	next = 'a'
	//	var accu = gp.NGramFreq[nGram+string(next)]
	//	fmt.Println("PickNGram 2: accu=", accu)
	//	for ; next <= 'z'; next++ {
	//		accu += gp.NGramFreq[nGram+string(next)]
	//		if accu > rnd {
	//			break
	//		}
	//	}
	//	fmt.Println("PickNGram 3: next=", string(next))
	//	isViablePrefix = (gp.NGramCnt[nGram[max(len(nGram)-1, 0):]+string(next)]) > 0
	//	fmt.Println("PickNGram 4: ", nGram[max(len(nGram)-1, 0):]+string(next), gp.NGramCnt[nGram[max(len(nGram)-1, 0):]+string(next)])
	//	count += 1
	//	if count > 100 {
	//		return gp.PickNGram(nGram[1:len(nGram)])
	//	}
	//}
	return next
}

func (gp GeneratorParams) Generate(length int) (password string) {
	if gp.maxNGramLen == 0 {
		gp.maxNGramLen = maxNGramLen(gp.NGramFreq)
	}
	if gp.maxNGramLen == 0 {
		return ""
	}
	password = ""
	for len(password) < length {
		password += string(gp.PickNGram(password[max(len(password)-gp.maxNGramLen, 0):len(password)]))
	}
	//password = string(gp.PickNGram(""))
	//fmt.Println("Generate 1", password)
	//password += string(gp.PickNGram(string(password[0])))
	//fmt.Println("Generate 2", password)
	//fmt.Println("Generate 3", gp.NGramFreq[password])
	//for len(password) < length {
	//	i := len(password)
	//	ch := string(gp.PickNGram(password[i-2 : i]))
	//	password += ch
	//}
	return password
}

//func Unravel(ngram *map[string]int, s *string) {
//	re, _ := regexp.Compile("([a-z]*)([0-9]*)")
//
//	for _, v := range re.FindAllStringSubmatch(*s, -1) {
//		(*ngram)[v[1]], _ = strconv.Atoi(v[2])
//	}
//}
