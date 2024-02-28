package recur

var permuteRes [][]int
var chosen []int
var chosenBool []bool
var chosenMap map[int]struct{}

func permute(nums []int) [][]int {
	permuteRes = nil
	chosenBool = make([]bool, len(nums))
	chosenMap = make(map[int]struct{})
	recurForPermute2(nums)
	return permuteRes
}

func recurForPermute(nums []int) {
	if len(chosen) == len(nums) {
		var tmpRes = make([]int, len(chosen))
		copy(tmpRes, chosen)
		permuteRes = append(permuteRes, tmpRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !chosenBool[i] {
			chosen = append(chosen, nums[i])
			chosenBool[i] = true
			recurForPermute(nums)
			chosen = chosen[:len(chosen)-1]
			chosenBool[i] = false
		}
	}
}

func recurForPermute2(nums []int) {
	if len(chosen) == len(nums) {
		var tmpRes = make([]int, len(chosen))
		copy(tmpRes, chosen)
		permuteRes = append(permuteRes, tmpRes)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, exist := chosenMap[i]; !exist {
			chosen = append(chosen, nums[i])
			chosenMap[i] = struct{}{}
			recurForPermute2(nums)
			chosen = chosen[:len(chosen)-1]
			delete(chosenMap, i)
		}
	}
}
