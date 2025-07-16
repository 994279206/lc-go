package leetcode

//删除排序数组中的重复项
//一个有序数组，原地删除重复出现的元素，使得每个元素只出现一次，返回新数组的长度
//不能使用额外的数组空间，必须在原地修改输入的数组并使用o(1)
//重点考察：双指针算法

func func002(a []int) int {
	if len(a) <= 0 {
		return 0
	}
	i := 0 //i：慢指针，j：快指针
	for j := 1; j < len(a); j++ {
		if a[i] != a[j] {
			i++
			a[i] = a[j]
		}
	}
	return i + 1
}
