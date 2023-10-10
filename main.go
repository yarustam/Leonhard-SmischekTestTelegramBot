package main

import (
	"fmt"
	"strings"
)

func countAnswer(myMap map[string][]int, answer string, i int) int {
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

func questionCycle(counts []int, QuestionsMaps []map[string][]int, arrTestQuestionList []string) {
	var answer string

	for i, question := range arrTestQuestionList {
		fmt.Printf("Вопрос №%d: %s \n", i+1, question)
		//ввел проверку на вводимый символ и перевод его в нижний регистр, для удобства
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
		for j, myMap := range QuestionsMaps {
			counts[j] += countAnswer(myMap, answer, i)
		}
	}
}

func main() {
	QuestionsMaps := []map[string][]int{
		{"д": {1, 11, 23, 33, 45, 55, 67, 77}},
		{"д": {9, 21, 43, 74, 87}, "н": {31, 53, 65}},
		{"д": {6, 18, 28, 40, 50, 62, 72, 84}},
		{"д": {8, 20, 30, 42, 52, 64, 75, 86}},
		{"д": {2, 15, 24, 34, 37, 56, 68, 78, 81}, "н": {12, 46, 59}},
		{"д": {3, 13, 35, 47, 57, 69, 79}, "н": {25}},
		{"д": {10, 32, 54, 76}},
		{"д": {6, 27, 38, 49, 60, 71, 82}, "н": {5}},
		{"д": {4, 14, 17, 26, 36, 48, 58, 61, 70, 80, 83}, "н": {39}},
		{"д": {7, 19, 22, 29, 41, 44, 63, 73, 85, 88}, "н": {51}},
	}

	counts := make([]int, 10)

	questionCycle(counts, QuestionsMaps, testQuestions)
	testResult(counts)

}
