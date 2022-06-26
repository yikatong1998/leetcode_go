package reverse_integer

import "math"

/**
执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
内存消耗：2 MB, 在所有 Go 提交中击败了 99.96% 的用户
通过测试用例：1032 / 1032
*/
func reverse2(x int) int {
	var ret int
	for x != 0 {
		if ret < math.MinInt32/10 || ret > math.MaxInt32/10 {
			return 0
		}

		n := x % 10
		x /= 10

		ret = ret*10 + n
	}

	return ret
}
