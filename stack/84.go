package stack

type Rec struct {
	Height int
	Width  int
}

type RecListNode struct {
	Val  *Rec
	Next *RecListNode
}

type RecListStack struct {
	Root *RecListNode
	Size int
}

func (l *RecListStack) Push(value *Rec) {
	if l.Root == nil {
		l.Root = &RecListNode{
			Val:  value,
			Next: nil,
		}
	} else {
		preRoot := l.Root
		l.Root = &RecListNode{
			Val:  value,
			Next: preRoot,
		}
	}

	l.Size++
}

func (l *RecListStack) Pop() *Rec {
	if l.Size == 0 {
		panic("empty")
	}
	top := l.Root.Val
	l.Root = l.Root.Next
	l.Size--
	return top
}

func (l *RecListStack) Peek() *Rec {
	if l.Size == 0 {
		panic("empty")
	}
	return l.Root.Val
}

func (l *RecListStack) Len() int {
	return l.Size
}

func (l *RecListStack) Empty() bool {
	return l.Size == 0
}

func largestRectangleArea(heights []int) int {
	var max = 0
	var recStack RecListStack
	for _, h := range heights {
		if recStack.Empty() {
			recStack.Push(&Rec{
				Height: h,
				Width:  1,
			})
		} else {
			top := recStack.Peek()
			var cul = 0
			for top.Height >= h {
				recStack.Pop()
				cul += top.Width
				if top.Height*cul > max {
					max = top.Height * cul
				}
				if recStack.Empty() {
					break
				}
				top = recStack.Peek()
			}
			recStack.Push(&Rec{
				Height: h,
				Width:  cul + 1, // 加上当前的宽
			})
		}
	}
	var cul = 0 // 统计整理后的栈
	for !recStack.Empty() {
		top := recStack.Pop()
		cul += top.Width
		if cul*top.Height > max {
			max = cul * top.Height
		}
	}
	return max
}

// 暴力法
func largestRectangleArea2(heights []int) int {
	var max = 0
	for index, h := range heights {
		var pre = 0
		i := index - 1
		for ; i >= 0; i-- {
			if heights[i] >= h {
				pre++
			} else {
				break
			}
		}
		var post = 0
		j := index + 1
		for ; j < len(heights); j++ {
			if heights[j] >= h {
				post++
			} else {
				break
			}
		}
		cur := h * (pre + post + 1)
		if cur > max {
			max = cur
		}
	}
	return max
}
