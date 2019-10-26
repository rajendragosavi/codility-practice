package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := int64(1255)
	a := strconv.FormatInt(n, 2)
	fmt.Println("binary form :-", a)
	//b := reflect.TypeOf(a)
	s := strings.Split(a, "")
	fmt.Println(s)
	count := 0
	countarray := make([]int, 0)
	for _, res := range s {

		if res == "1" {
			countarray = append(countarray, count)
			count = 0
		} else {
			count++
		}

	}
	fmt.Println("count array :- ", countarray)
	max := countarray[0]
	for _, j := range countarray {
		if j > max {
			max = j
		}
	}
	fmt.Println("largest gap is :-", max)
}
