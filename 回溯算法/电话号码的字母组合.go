var str string
var rets = make([]string, 0)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	// 构造数据
	d2a := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	digs := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		n := string(digits[i])
		digs = append(digs, n)
	}
	get(d2a, digs, 0)

	tmp := make([]string, len(rets))
	copy(tmp, rets)
	// 清空， 方法会被多次调用
	rets = rets[:0]
	str = ""
	return tmp
}

func get(d2a map[string][]string, digs []string, digsindex int) {
	digsize := len(digs)
	src := d2a[digs[digsindex]]

	for j := 0; j < len(src); j++ {
		str += src[j]
		if digsindex+1 < digsize {
			// 递归调用
			get(d2a, digs, digsindex+1)
			// 回溯
			str = str[:digsindex]
		}
		if len(str) == digsize {
			rets = append(rets, str)
			// 回溯
			str = str[:digsindex]
		}
	}
}
