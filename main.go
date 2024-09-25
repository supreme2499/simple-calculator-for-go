package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	arab = map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	rim = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
)

func main() {

	var tns string
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	arr := strings.Split(str, " ")

	// проверка размера строки
	if len(arr) > 3 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	if len(arr) < 3 {
		panic("строка не является математической операцией")
	}

	//Для удобства работы заносим числа в массив, а оператор в переменную
	nums := []string{arr[0], arr[2]}
	oper := arr[1]

	// проверка корректности систем счисления
	if status(nums[0]) == status(nums[1]) {
		tns = status(nums[0])
	} else {
		panic("используются одновременно разные системы счисления")
	}

	if tns == "arab" {
		fmt.Println(mathem(arab, oper, nums))
	} else {
		num := mathem(rim, oper, nums)
		if num < 1 {
			panic("римской системе нет нуля и отрицательных чисел")
		}
		fmt.Println(convert(num))

	}

}

// делаем проверку того, какие числа нам дали, в какой системе счисления
func status(nums string) (status string) {
	_, okr := rim[nums]
	if okr {
		status = "rim"
	}
	_, oka := arab[nums]
	if oka {
		status = "arab"
	}
	return status
}

func mathem(sys map[string]int, oper string, nums []string) int {
	switch oper {
	case "+":
		return (sys[nums[0]] + sys[nums[1]])
	case "-":
		return (sys[nums[0]] - sys[nums[1]])
	case "/":
		return (sys[nums[0]] / sys[nums[1]])
	case "*":
		return (sys[nums[0]] * sys[nums[1]])
	default:
		panic("используется неверный оператор")
	}
}

func convert(num int) string {
	var builder strings.Builder

	rim := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for i := 0; num > 0; {
		if rim[i].Value <= num {
			builder.WriteString(rim[i].Symbol)
			num -= rim[i].Value
			i = 0
			continue
		}
		i++
	}

	return builder.String()
}
