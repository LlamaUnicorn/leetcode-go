package mco

var nums = []int{1, 1, 0, 1, 1, 1}

func MaxConsecutiveOnes() int {
	startPtr := 0
	//endPtr := 0
	for i := 0; i < len(nums); {
		if nums[i] == 1 {
			startPtr = i
		}
		return startPtr
	}
	return len(nums)
}
