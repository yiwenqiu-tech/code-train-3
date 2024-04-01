package greedy

func lemonadeChange(bills []int) bool {
	moneyMap := make(map[int]int)
	for _, bill := range bills {
		moneyMap[bill]++
		if !canChange(bill, moneyMap) {
			return false
		}
	}
	return true
}

func canChange(bill int, moneyMap map[int]int) bool {
	needChange := bill - 5
	if needChange == 0 {
		return true
	}
	for _, coin := range []int{20, 10, 5} {
		for needChange >= coin {
			if value, exist := moneyMap[coin]; exist && value > 0 {
				moneyMap[coin]--
				needChange -= coin
			} else {
				break
			}
		}
	}
	return needChange == 0
}
