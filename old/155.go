package main

//MinStack stack存的值為該值與min的差距，用min來存目前的最小值
type MinStack struct {
	stack []int
	min   int
}

//Constructor 建構子
func Constructor() MinStack {
	stack := MinStack{[]int{}, 0}
	return stack
}

//Push add element
func (stack *MinStack) Push(x int) {
	if 0 == len(stack.stack) {
		//存入stack的第一個值，min為此值、gap為0
		stack.stack = append(stack.stack, 0)
		stack.min = x
	} else {
		//存入gap，若傳入的值比目前的min還小則需更改min值
		stack.stack = append(stack.stack, x-stack.min)
		if x < stack.min {
			stack.min = x
		}
	}
}

//Pop delete top element
func (stack *MinStack) Pop() {
	l := len(stack.stack)
	val := stack.stack[l-1] //取得top gap
	stack.stack = stack.stack[:l-1]

	if val < 0 {
		// top gap < 0，代表該直輸入時比min小，有因此調整過min
		// 所以在此須將min再次調整回來。
		stack.min -= val
	}
}

//Top get the top value
func (stack *MinStack) Top() int {
	val := stack.stack[len(stack.stack)-1]

	if val > 0 {
		//gap > 0，代表輸入值比當前min還大，因此回傳值需調整
		return val + stack.min
	}
	//gap < 0的狀況代表輸入時已調整過min值，因此回傳min值即可
	//gap = 0 表示該值與min值相等
	return stack.min
}

//GetMin return current min value
func (stack *MinStack) GetMin() int {
	return stack.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
