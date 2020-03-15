package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
}

func dayOfTheWeek(day int, month int, year int) string {
	layout := "2/1/2006"
	date := strconv.Itoa(day) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(year)
	t, _ := time.Parse(layout, date)
	return t.Weekday().String()
}
