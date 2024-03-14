package binary_search

type TopVotedCandidate struct {
	Persons         []int
	Times           []int
	TopVotedForTime map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	topVotedForTime := make(map[int]int) // 某个时刻最高票数
	topVoted := 0
	lastMaxVotedPerson := 0
	personVoted := make(map[int]int)
	for index, t := range times {
		var curVoted = 0
		curVoted, _ = personVoted[persons[index]]
		if curVoted+1 >= topVoted { // 当前票数>=最高票，则当前人为此时刻的最高票者
			topVoted = curVoted + 1
			topVotedForTime[t] = persons[index]
			lastMaxVotedPerson = persons[index]
		} else { // 否则还是上次最高票者
			topVotedForTime[t] = lastMaxVotedPerson
		}
		personVoted[persons[index]]++
	}
	return TopVotedCandidate{
		Persons:         persons,
		Times:           times,
		TopVotedForTime: topVotedForTime,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	left := 0
	right := len(this.Times) - 1
	for left < right {
		mid := left + (right-left+1)/2
		if this.Times[mid] <= t {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return this.TopVotedForTime[this.Times[right]]
}
