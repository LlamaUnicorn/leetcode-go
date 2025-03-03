package main

var arr = []int{1, 0, 2, 3, 0, 4, 5, 0} // [1,0,0,2,3,0,0,4]

func main() {
	duplicateZeros(arr)
}

func duplicateZeros(arr []int) {
	zeroes := 0

	for _, v := range arr {
		if v == 0 {
			zeroes++
		}
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if zeroes+i < len(arr) {
				arr[zeroes+i] = 0
			}

			if zeroes-1+i < len(arr) {
				arr[zeroes-1+i] = 0
			}

			zeroes--
		} else if i+zeroes < len(arr) {
			arr[zeroes+i] = arr[i]
		}
	}
}
