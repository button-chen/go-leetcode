type node struct {
	word string
	subs []*node
}

type MagicDictionary struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	if this.root == nil {
		this.root = new(node)
	}
	for _, word := range dict {
		tmp := this.root
		for i := 0; i < len(word); i++ {
			if tmp.subs == nil {
				tmp.subs = make([]*node, 26)
			}
			code := int(word[i]) - int('a')
			if tmp.subs[code] == nil {
				tmp.subs[code] = new(node)
			}
			tmp = tmp.subs[code]
		}
		tmp.word = word
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	return this.search(this.root, word)
}

func (this *MagicDictionary) find(root *node, word string) bool {
	tmp := root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			return false
		}
		code := int(word[i]) - 97
		if tmp.subs[code] != nil {
			tmp = tmp.subs[code]
		} else {
			return false
		}
		if i == len(word)-1 {
			if tmp.word != "" {
				return true
			}
		}
	}
	return false
}

func (this *MagicDictionary) search(root *node, word string) bool {
	tmp := root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			return false
		}
		code := int(word[i]) - 97
		tmpcode := code
		if tmp.subs[code] == nil {
			tmpcode = -1
		}
		for k, p := range tmp.subs {
			if p == nil {
				continue
			}
			if k == tmpcode {
				continue
			}
			if i == len(word)-1 {
				if p.word != "" {
					return true
				}
				return false
			}
			if this.find(p, word[i+1:]) {
				return true
			}
		}
		if tmp.subs[code] != nil {
			tmp = tmp.subs[code]
		} else {
			return false
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
