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
		fmt.Println(mathem(oper, nums))
	} else {
		num := mathem(oper, nums)
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

func mathem(oper string, nums []string) int {
	switch oper {
	case "+":
		return (arab[nums[0]] + arab[nums[1]])
	case "-":
		return (arab[nums[0]] - arab[nums[1]])
	case "/":
		return (arab[nums[0]] / arab[nums[1]])
	case "*":
		return (arab[nums[0]] * arab[nums[1]])
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
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}

	for i := 0; num > 0; {

		if rim[i].Value <= num {
			builder.WriteString(rim[i].Symbol)
			num -= rim[i].Value
			continue
		} else {
			for j := i + 1; j < len(rim); j++ {
				if rim[i].Value-rim[j].Value <= num {
					if rim[i].Value-rim[j].Value == 5 {
						continue
					}
					builder.WriteString(rim[j].Symbol)
					builder.WriteString(rim[i].Symbol)
					num -= rim[i].Value - rim[j].Value
					i = 0
					break
				}
			}
		}
		i++
	}

	return builder.String()
}
