package main

import "fmt"

func separator(array ...int) (even []int, odd []int) { // если правильно понял, то под массивами имелись в виду срезы

	for _, v := range array {

		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return
}

func main() {

	fmt.Println("23.6 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Чётные и нечётные.")
	fmt.Println("------------")

	even, odd := separator(1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(even)
	fmt.Println(odd)
}
