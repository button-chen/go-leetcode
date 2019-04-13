import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var vals = make([]int, 0, 100)

func getNums(root *TreeNode) {
	if root != nil {
		vals = append(vals, root.Val)
		getNums(root.Left)
		getNums(root.Right)
	}
}

func minDiffInBST(root *TreeNode) int {
	getNums(root)
	sort.Ints(vals)

	minval := math.MaxInt32
	for i := 0; i < len(vals)-1; i++ {
		tmp := int(math.Abs(float64(vals[i+1] - vals[i])))
		if tmp < minval {
			minval = tmp
			if minval <= 1 {
				break
			}
		}
	}
	// 必须清理，自动测试会多次调用minDiffInBST
	vals = vals[:0]
	return minval
}
  