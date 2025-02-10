package bm1342

//func NumberOfSteps(num int) int {
//	count := 0
//	for i := num; i > 0; {
//		if i%2 == 0 {
//			i /= 2
//			count++
//		} else {
//			i -= 1
//			count++
//		}
//	}
//
//	return count
//}

func NumberOfSteps(num int) int {
	if num == 0 {
		return 0 // Edge case: no steps needed for 0
	}

	// Count the number of 1s in the binary representation
	onesCount := 0
	for n := num; n > 0; n >>= 1 {
		if n&1 == 1 {
			onesCount++
		}
	}

	// Total bits in binary representation
	totalBits := 0
	for n := num; n > 0; n >>= 1 {
		totalBits++
	}

	// Total steps = number of 1s + (total bits - 1)
	return onesCount + (totalBits - 1)
}

//func NumberOfSteps(num int) int {
//	if num == 0 {
//		return 0
//	}
//
//	steps := 0
//	for num > 0 {
//		if num&1 == 1 { // If the number is odd
//			steps++ // Subtract 1
//		}
//		if num > 1 { // If the number is greater than 1, divide by 2
//			steps++ // Right shift
//		}
//		num >>= 1 // Perform the division (right shift)
//	}
//	return steps
//}
