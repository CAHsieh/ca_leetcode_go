package main

func sortColors(nums []int) {
	redC := 0
	whiteC := 0
	for _, v := range nums {
		if v == 0 {
			redC++
		} else if v == 1 {
			whiteC++
		}
	}

	whiteC += redC

	for i := range nums {
		if i < redC {
			nums[i] = 0
		} else if i < whiteC {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
