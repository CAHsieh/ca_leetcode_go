package main

import (
	"fmt"
	"sort"
)

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}

type KthLargest struct {
	k      int
	arr    []int
	length int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k, nums, len(nums)}
}

func (this *KthLargest) Add(val int) int {
	if this.length < this.k {
		this.arr = append(this.arr, val)
		for i := this.length - 1; i >= 0; i-- {
			if val >= this.arr[i] || i == 0 {
				if i == 0 && val < this.arr[i] {
					i--
				}
				this.arr = append(this.arr[:i+1], append([]int{val}, this.arr[i+1:]...)...)
				break
			}
		}
	} else if val > this.arr[this.length-this.k] {
		for i := this.length - this.k; i < this.length; i++ {
			if val < this.arr[i] || i == this.length-1 {
				if i == this.length-1 && val > this.arr[i] {
					i++
				}
				this.arr = append(this.arr[:i], append([]int{val}, this.arr[i:]...)...)
				break
			}
		}
	} else {
		for i := this.length - this.k; i >= 0; i-- {
			if val >= this.arr[i] || i == 0 {
				if i == 0 && val < this.arr[i] {
					i--
				}
				this.arr = append(this.arr[:i+1], append([]int{val}, this.arr[i+1:]...)...)
				break
			}
		}
	}

	this.length++
	return this.arr[this.length-this.k]
}
