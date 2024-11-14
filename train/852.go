package train

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		lMid := left + (right-left)/2
		rMid := lMid + 1
		if arr[lMid] > arr[rMid] {
			right = rMid - 1
		} else {
			left = lMid + 1
		}
	}
	return right
}
