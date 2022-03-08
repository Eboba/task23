package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences [4]string, chars [5]rune) (output [4][5]int) {

	var endRune int

	for i, v := range sentences {
		endRune = 0
		a := []rune(v)
		b := chars
		for len(a)-1 >= 0 {
			for j, w := range b {
				if strings.EqualFold(string(a[len(a)-1:len(a)]), string(w)) { // Сравнивает строки без учета регистра
					output[i][j] = len(a)
					b[j] = '/' // Убераем руны которые уже нашлись
					endRune++
				}
			}
			if endRune == len(chars) { //Останавливает цикл, когда кончились искомые руны
				break
			}
			a = a[:len(a)-1]
		}
	}
	return
}

func main() {

	fmt.Println("23.6 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 2. Поиск символов в нескольких строках.")
	fmt.Println("------------")

	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	var array [4][5]int
	array = parseTest(sentences, chars)
	for i := 0; i < len(array); i++ {
		fmt.Println("Предложение:", i+1)
		for j := 0; j < len(chars); j++ {
			if array[i][j] == 0 {
				continue
			} else {
				fmt.Printf("'%v' position %v\n", string(chars[j]), array[i][j]-1)
			}
		}
		fmt.Println()
	}
}
