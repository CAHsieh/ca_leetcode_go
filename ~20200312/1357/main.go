package main

import "fmt"

func main() {
	cashier := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	fmt.Println(cashier.GetBill([]int{1, 2}, []int{1, 2}))                                  // return 500.0, bill = 1 * 100 + 2 * 200 = 500.
	fmt.Println(cashier.GetBill([]int{3, 7}, []int{10, 10}))                                // return 4000.0
	fmt.Println(cashier.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}))    // return 800.0, The bill was 1600.0 but as this is the third customer, he has a discount of 50% which means his bill is only 1600 - 1600 * (50 / 100) = 800.
	fmt.Println(cashier.GetBill([]int{4}, []int{10}))                                       // return 4000.0
	fmt.Println(cashier.GetBill([]int{7, 3}, []int{10, 10}))                                // return 4000.0
	fmt.Println(cashier.GetBill([]int{7, 5, 3, 1, 6, 4, 2}, []int{10, 10, 10, 9, 9, 9, 7})) // return 7350.0, Bill was 14700.0 but as the system counted three more customers, he will have a 50% discount and the bill becomes 7350.0
	fmt.Println(cashier.GetBill([]int{2, 3, 5}, []int{5, 3, 2}))                            // return 2500.0
}

type Cashier struct {
	n        int
	countN   int
	discount int
	products map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	maps := make(map[int]int)
	for i := 0; i < len(products); i++ {
		maps[products[i]] = prices[i]
	}
	return Cashier{n, n, discount, maps}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.countN--
	length := len(product)
	sum := 0
	for i := 0; i < length; i++ {
		sum += (this.products[product[i]] * amount[i])
	}
	var total float64 = float64(sum)
	if this.countN == 0 {
		//todo discount
		total = total - (total * (float64(this.discount) / 100))
		this.countN = this.n
	}
	return total
}
