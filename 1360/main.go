package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
}

func daysBetweenDates(date1 string, date2 string) int {

	arr := strings.Split(date1, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])
	time1 := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	arr = strings.Split(date2, "-")
	year, _ = strconv.Atoi(arr[0])
	month, _ = strconv.Atoi(arr[1])
	day, _ = strconv.Atoi(arr[2])

	time2 := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	duration := time2.Sub(time1)

	return int(math.Abs((duration.Hours() / 24)))
}
