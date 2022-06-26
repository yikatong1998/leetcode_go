package solution

/**
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给你一个整数，将其转为罗马数字。



示例 1:

输入: num = 3
输出: "III"
示例 2:

输入: num = 4
输出: "IV"
示例 3:

输入: num = 9
输出: "IX"
示例 4:

输入: num = 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: num = 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/

/**
执行用时：4 ms, 在所有 Go 提交中击败了 92.85% 的用户
内存消耗：3.3 MB, 在所有 Go 提交中击败了 18.66% 的用户
通过测试用例：3999 / 3999
*/
func intToRoman(num int) string {
	if num > 3999 {
		return ""
	}

	var ret string
	var r []int
	for count := 0; num > 0; count++ {
		r = append(r, num%10)
		num /= 10
	}

	for i, n := range r {
		if n == 0 {
			continue
		}

		var base int
		switch i {
		case 0:
			base = 1
		case 1:
			base = 10
		case 2:
			base = 100
		case 3:
			base = 1000
		default:
			base = 1
		}

		roman := getRoman(n, base)
		ret = roman + ret
	}

	return ret
}

func getRoman(n, base int) string {
	var ret string

	sign := false
	if n == 4 || n == 5 || n == 9 {
		sign = true
	}

	if n == 6 || n == 7 || n == 8 {
		return getRoman(5, base) + getRoman(n-5, base)
	}

	m := n * base
	if !sign {
		m = base
	}

	switch m {
	case 1:
		ret = "I"
	case 4:
		ret = "IV"
	case 5:
		ret = "V"
	case 9:
		ret = "IX"
	case 10:
		ret = "X"
	case 40:
		ret = "XL"
	case 50:
		ret = "L"
	case 90:
		ret = "XC"
	case 100:
		ret = "C"
	case 400:
		ret = "CD"
	case 500:
		ret = "D"
	case 900:
		ret = "CM"
	case 1000:
		ret = "M"
	}

	if !sign {
		s := ret
		for i := 0; i < n-1; i++ {
			ret += s
		}
	}

	return ret
}
