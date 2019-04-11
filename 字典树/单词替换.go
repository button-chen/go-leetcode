import "strings"

type node struct {
	word string
	subs []*node
}

// 返回找到的最短词根
func find(root *node, word string) (string, bool) {
	tmp := root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			break
		}
		code := int(word[i]) - 97
		if tmp.subs[code] != nil {
			tmp = tmp.subs[code]
		} else {
			break
		}
		if tmp.word != "" {
			return tmp.word, true
		}
	}
	return "", false
}

func replaceWords(dict []string, sentence string) string {
	// create trie
	root := new(node)
	root.subs = make([]*node, 26)
	for _, word := range dict {
		tmp := root
		for i := 0; i < len(word); i++ {
			if tmp.subs == nil {
				tmp.subs = make([]*node, 26)
			}
			code := int(word[i]) - 97
			if tmp.subs[code] == nil {
				tmp.subs[code] = new(node)
			}
			tmp = tmp.subs[code]
		}
		tmp.word = word
	}
	// search
	ret := make([]string, 0)
	sentenses := strings.Split(sentence, " ")
	for i := 0; i < len(sentenses); i++ {
		str, ok := find(root, sentenses[i])
		if !ok {
			ret = append(ret, sentenses[i])
		} else {
			ret = append(ret, str)
		}
	}
	return strings.Join(ret, " ")
}
