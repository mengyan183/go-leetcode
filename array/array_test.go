package array

import (
	"sort"
	"testing"
)

//https://leetcode-cn.com/problems/3sum
func TestThreeNum(t *testing.T) {
	var nums []int = []int{1, 2, -2, -1}
	sum := threeSumCustom(nums)
	t.Log(sum)
}
// 这里实际利用了排序和快慢指针进行操作
// 对于排序直接使用现成的排序工具; 之所以使用排序的原因在于为了满足出现的结果不可重复,当进行排序后,就可以保证相同的数据肯定是紧邻的,因此我们只需要对比紧邻数据是否相同,如果不相同则肯定不会出现相同的结果
func threeSumCustom(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	l := len(nums)
	for left := 0; left < len(nums); left++ {
		// 保证不会出现重复的数据
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		right := l - 1
		// 第二个指针
		for m := left + 1; m < len(nums); m++ {
			if m > left+1 && nums[m] == nums[m-1] {
				continue
			}
			if right <= m {
				break
			}
			// 第三个指针
			//right := l - 1
			// 对于right指针操作,从尾部向前遍历,当不存在三个指针对应的值相加大于0时,由于数据本身时递增的,因此在right左侧和m右侧的数据的结果则有可能出现为0的数据
			// 对于 right声明在第二个循环外部的原因,结束以下循环的条件为 三者值相加结果小于等于0,而m是递增的,对应的结果只可能大于上一次的值,因此对于当前right指针
			// 假设 m 对应的值为 a, m+1对应的值为a' ;right 对应的值为 b ; right+1 对应的值为 b' , left对应的值为c
			// 对于结束循环的条件有可能存在 a+b+c <=0 则肯定存在 a+b'+c >0 ; b' - b >0 ; a' - a >0 ;因此对于 a+b'+c+a'-a = a'+b'+c >0是必然的结果;而对于 a'+b+c的结果不确定
			// 所以对于下一次循环操作只需要从当前right位置开始遍历即可,对于right后的数据无需再次进行遍历,避免无用计算
			for ; right > m && nums[left]+nums[m]+nums[right] > 0; right-- {

			}
			// 对于上述for循环有可能会出现right == m 的情况,因此需要排除
			if right == m {
				break
			}
			if nums[left]+nums[m]+nums[right] == 0 {
				result = append(result, []int{nums[left], nums[m], nums[right]})
			}
		}
	}
	return result
}
