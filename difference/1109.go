package difference

func corpFlightBookings(bookings [][]int, n int) []int {
	var diffArray = make([]int, n+2)
	for _, booking := range bookings {
		first := booking[0]
		last := booking[1]
		count := booking[2]

		diffArray[first] += count
		diffArray[last+1] -= count
	}
	var res = make([]int, n+2)
	res[0] = diffArray[0]
	for i := 1; i < len(diffArray); i++ {
		res[i] = res[i-1] + diffArray[i]
	}
	return res[1 : n+1]
}
