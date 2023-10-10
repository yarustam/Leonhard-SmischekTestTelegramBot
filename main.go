package main

import (
	"fmt"
	"strings"
)

func littleCycle(myMap map[string][]int, answer string, i int) int {
	count := 0
	if nums, ok := myMap[answer]; ok {
		for _, num := range nums {
			if num == i+1 {
				count++
				break
			}
		}
	}
	return count
}

func questionCycle(countA, countB, countC, countD, countE, countF, countG, countH, countI, countJ *int, myMapA, myMapB, myMapC, myMapD, myMapE, myMapF, myMapG, myMapH, myMapI, myMapJ map[string][]int, testQuestions []string) {
	var answer string

	for i, question := range testQuestions {
		fmt.Printf("Вопрос №%d: %s \n", i+1, question)
		for {
			fmt.Print("Ваш ответ(д/н): ")
			fmt.Scan(&answer)
			answer = strings.ToLower(answer)

			if answer != "д" && answer != "н" {
				fmt.Println("Пожалуйста, отвечайте д/н.")
				continue
			}
			break
		}

		*countA += littleCycle(myMapA, answer, i)
		*countB += littleCycle(myMapB, answer, i)
		*countC += littleCycle(myMapC, answer, i)
		*countD += littleCycle(myMapD, answer, i)
		*countE += littleCycle(myMapE, answer, i)
		*countF += littleCycle(myMapF, answer, i)
		*countG += littleCycle(myMapG, answer, i)
		*countH += littleCycle(myMapH, answer, i)
		*countI += littleCycle(myMapI, answer, i)
		*countJ += littleCycle(myMapJ, answer, i)
	}
}

func testResult(countA, countB, countC, countD, countE, countF, countG, countH, countI, countJ int) {
	fmt.Printf("Количество совпадений в группе A: %d\n", countA)
	fmt.Printf("Количество совпадений в группе B: %d\n", countB)
	fmt.Printf("Количество совпадений в группе C: %d\n", countC)
	fmt.Printf("Количество совпадений в группе D: %d\n", countD)
	fmt.Printf("Количество совпадений в группе E: %d\n", countE)
	fmt.Printf("Количество совпадений в группе F: %d\n", countF)
	fmt.Printf("Количество совпадений в группе G: %d\n", countG)
	fmt.Printf("Количество совпадений в группе H: %d\n", countH)
	fmt.Printf("Количество совпадений в группе I: %d\n", countI)
	fmt.Printf("Количество совпадений в группе J: %d\n", countJ)
}

func main() {

	myMapA := make(map[string][]int)
	myMapB := make(map[string][]int)
	myMapC := make(map[string][]int)
	myMapD := make(map[string][]int)
	myMapE := make(map[string][]int)
	myMapF := make(map[string][]int)
	myMapG := make(map[string][]int)
	myMapH := make(map[string][]int)
	myMapI := make(map[string][]int)
	myMapJ := make(map[string][]int)

	myMapA["д"] = []int{1, 11, 23, 33, 45, 55, 67, 77}

	myMapB["д"] = []int{9, 21, 43, 74, 87}
	myMapB["н"] = []int{31, 53, 65}

	myMapC["д"] = []int{6, 18, 28, 40, 50, 62, 72, 84}

	myMapD["д"] = []int{8, 20, 30, 42, 52, 64, 75, 86}

	myMapE["д"] = []int{2, 15, 24, 34, 37, 56, 68, 78, 81}
	myMapE["н"] = []int{12, 46, 59}

	myMapF["д"] = []int{3, 13, 35, 47, 57, 69, 79}
	myMapF["н"] = []int{25}

	myMapG["д"] = []int{10, 32, 54, 76}

	myMapH["д"] = []int{6, 27, 38, 49, 60, 71, 82}
	myMapH["н"] = []int{5}

	myMapI["д"] = []int{4, 14, 17, 26, 36, 48, 58, 61, 70, 80, 83}
	myMapI["н"] = []int{39}

	myMapJ["д"] = []int{7, 19, 22, 29, 41, 44, 63, 73, 85, 88}
	myMapJ["н"] = []int{51}

	countA := 0
	countB := 0
	countC := 0
	countD := 0
	countE := 0
	countF := 0
	countG := 0
	countH := 0
	countI := 0
	countJ := 0

	questionCycle(&countA, &countB, &countC, &countD, &countE, &countF, &countG, &countH, &countI, &countJ, myMapA, myMapB, myMapC, myMapD, myMapE, myMapF, myMapG, myMapH, myMapI, myMapJ, testQuestions)
	testResult(countA, countB, countC, countD, countE, countF, countG, countH, countI, countJ)

}
