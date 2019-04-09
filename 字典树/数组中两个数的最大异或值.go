type node struct {
	val  int
	subs []*node
}

func findMaximumXOR(nums []int) int {
	// 构建trie
	root := new(node)
	for _, v := range nums {
		tmp := root
		for i := 31; i >= 0; i-- {
			if tmp.subs == nil {
				tmp.subs = make([]*node, 2)
			}
			bit := (v >> uint(i)) & 1
			if tmp.subs[bit] == nil {
				tmp.subs[bit] = new(node)
			}
			tmp = tmp.subs[bit]
		}
	}
	// search
	// 基于最高位为1大于后面任意的值， 例如:  10000000 大于 01111111
	// 基于异或特点: a ^ b = c  a ^ c = b
	maxval := -1
	for _, v := range nums {
		tmp := root
		tmpmax := 0
		for i := 31; i >= 0; i-- {
			bit := (v >> uint(i)) & 1
			// 存在一个分支 ^ bit = 1
			if tmp.subs[bit^1] != nil {
				tmpmax += 1 << uint(i)
				tmp = tmp.subs[bit^1]
			} else {
				tmp = tmp.subs[bit]
			}
		}
		if tmpmax > maxval {
			maxval = tmpmax
		}
	}
	return maxval
}
