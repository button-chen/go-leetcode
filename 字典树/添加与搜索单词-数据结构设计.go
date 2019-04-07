type node struct {
	word string
	subs []*node
}

type WordDictionary struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if this.root == nil {
		this.root = new(node)
	}
	tmp := this.root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			tmp.subs = make([]*node, 26)
		}
		code := int(word[i]) - 97
		if tmp.subs[code] == nil {
			tmp.subs[code] = new(node)
		}

		tmp = tmp.subs[code]
		if i == len(word)-1 {
			tmp.word = word
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word)
}

func (this *WordDictionary) search(root *node, word string) bool {
	if root == nil {
		return false
	}
	tmp := root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			return false
		}
		if int(word[i]) == int('.') {
			for k, p := range tmp.subs {
				if p == nil {
					continue
				}
				if i == len(word)-1 {
					if p.word != "" {
						return true
					} else {
						if k == len(tmp.subs)-1 {
							return false
						}
						continue
					}
				}
				if this.search(p, word[i+1:]) {
					return true
				}
			}
			return false
		}
		code := int(word[i]) - 97
		if tmp.subs[code] == nil {
			return false
		}

		tmp = tmp.subs[code]
		if i == len(word)-1 {
			if tmp.word != "" {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
 