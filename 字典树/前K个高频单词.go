import "sort"

type node struct {
	mark bool
	subs []*node
	cnt  int
}

func find(root *node, word string) int {
	tmp := root
	for i := 0; i < len(word); i++ {
		if tmp.subs == nil {
			return -1
		}
		code := int(word[i]) - 97
		if tmp.subs[code] == nil {
			return -1
		}
		tmp = tmp.subs[code]
		if i == len(word)-1 {
			if tmp.mark {
				tmp.mark = false
				return tmp.cnt
			}
			break
		}
	}
	return -1
}

func topKFrequent(words []string, k int) []string {
	// create trie
	root := new(node)
	for _, word := range words {
		tmp := root
		for i := 0; i < len(word); i++ {
			code := int(word[i]) - 97
			if tmp.subs == nil {
				tmp.subs = make([]*node, 26)
			}
			if tmp.subs[code] == nil {
				tmp.subs[code] = new(node)
			}
			tmp = tmp.subs[code]
		}
		tmp.mark = true
		tmp.cnt++
	}

	// search
	ret := make(map[int][]string)
	for _, word := range words {
		cnt := find(root, word)
		if cnt != -1 {
			if _, ok := ret[cnt]; !ok {
				ret[cnt] = make([]string, 0)
			}
			ret[cnt] = append(ret[cnt], word)
		}
	}
	idx := make([]int, 0)
	for k := range ret {
		idx = append(idx, k)
	}
	sort.Ints(idx)
	// fetch
	result := make([]string, 0)
	size := len(idx) - 1
	for i := size; i >= 0; i-- {
		tmp := ret[idx[i]]
		sort.Strings(tmp)
		result = append(result, tmp...)
		if len(result) >= k {
			break
		}
	}
	return result[0:k]
}
