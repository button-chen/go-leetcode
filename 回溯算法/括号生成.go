func generateParenthesis(n int) []string {
	vret := make([]string, 0)
	get(n, n, &vret, "")
	return vret
}

// left:  左括号("(") 当前剩余个数
// right: 右括号(")") 当前剩余个数
func get(left, right int, vret *[]string, str string) {
	// 左括号剩余的个数应该小与右括号剩余的个数
	if left > right {
		return
	}
	if right == 0 && left == 0 {
		*vret = append(*vret, str)
	}
	if left > 0 {
		get(left-1, right, vret, str+"(")
	}
	if right > 0 {
		get(left, right-1, vret, str+")")
	}
}
