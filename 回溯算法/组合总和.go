import "sort"

var vret = make([][]int, 0)
var sums = make([]int, 0)
var sum int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	get(candidates, target, 0)
	// 会被多次调用， 清理全局状态
	tmp := vret
	vret = make([][]int, 0)
	sums = sums[:0]
	sum = 0
	return tmp
}

func get(candidates []int, target int, start int) {
	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		sums = append(sums, candidates[i])

		if sum < target {
			// start 开始位置参数实现去重复功效
			get(candidates, target, i)
		}
		if sum == target {
			tmp := make([]int, 0)
			tmp = append(tmp, sums...)
			vret = append(vret, tmp)
		}
		// 回溯
		sums = sums[:len(sums)-1]
		sum -= candidates[i]
		if sum > target {
			// 优化: 因为是排序的，往后走也一样是大于，所以不用继续
			break
		}
	}
}
