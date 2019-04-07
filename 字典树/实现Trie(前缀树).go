type node struct {
	word string
	subs []*node
}

type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
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

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this.root == nil {
		return false
	}
	tmp := this.root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			return false
		}
		code := int(word[i]) - 97
		if tmp.subs[code] == nil {
			return false
		}
		tmp = tmp.subs[code]
		if i == len(word)-1 {
			if tmp.word == word {
				return true
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this.root == nil {
		return false
	}
	tmp := this.root
	for i := 0; i < len(prefix); i++ {
		if tmp.subs == nil {
			tmp.subs = make([]*node, 26)
		}
		code := int(prefix[i]) - 97
		if tmp.subs[code] == nil {
			return false
		}
		tmp = tmp.subs[code]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
 