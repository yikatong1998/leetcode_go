package reverse_integer

import "math"

/**
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-2^31 <= x <= 2^31 - 1
*/

/**
执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了 5.16% 的用户
通过测试用例：1032 / 1032
*/

func reverse(x int) int {
	sign := false
	if x < 0 {
		sign = true
	}

	maxVal := uint32(math.MaxInt32)
	if sign {
		maxVal++
	}

	maxBits := getBits(maxVal)
	start, end := 0, len(maxBits)-1
	for start < end {
		maxBits[start], maxBits[end] = maxBits[end], maxBits[start]
		start++
		end--
	}

	var newX uint32
	if x == math.MinInt32 {
		newX = uint32(x)
	} else if x < 0 {
		newX = uint32(-x)
	} else {
		newX = uint32(x)
	}

	xBits := getBits(newX)
	if isBigger(xBits, maxBits) {
		return 0
	}

	ret := getNum(xBits)
	if sign {
		ret = -ret
	}

	return ret
}

func getBits(num uint32) []int {
	var ret []int
	for num > 0 {
		n := num % 10
		ret = append(ret, int(n))
		num /= 10
	}

	return ret
}

func getNum(a []int) int {
	var ret int
	for i := len(a) - 1; i >= 0; i-- {
		ret += a[i] * int(math.Pow10(len(a)-1-i))
	}

	return ret
}

func isBigger(a, b []int) bool {
	if len(a) > len(b) {
		return true
	} else if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		na, nb := a[i], b[i]

		if na > nb {
			return true
		} else if na < nb {
			return false
		}
	}

	return false
}
