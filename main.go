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

func questionCycle(countA, countB, countC *int, myMapA, myMapB, myMapC map[string][]int, testQuestions []string) {
	var answer string

	for i, question := range testQuestions {
		fmt.Printf("Вопрос №%d: %s \n", i+1, question)
		
		for {
			fmt.Print("Введите 'д' или 'н': ")
			fmt.Scan(&answer)
			answer = strings.ToLower(answer)
			if answer != "д" && answer != "н" {
				fmt.Println("Некорректный ввод. Введите 'д' или 'н'.")
				continue
			}
			break
		}

		*countA += littleCycle(myMapA, answer, i)
		*countB += littleCycle(myMapB, answer, i)
		*countC += littleCycle(myMapC, answer, i)
	}
}

func testResult(countA, countB, countC int) {
    fmt.Printf("Количество совпадений в группе A: %d\n", countA)
    fmt.Printf("Количество совпадений в группе B: %d\n", countB)
    fmt.Printf("Количество совпадений в группе C: %d\n", countC)
}

func main() {
    
    testQuestions  := []string{
  	  "У вас чаще веселое и беззаботное настроение?",
      "Вы чувствительны к оскорблениям?",
      "Бывает ли так, что у вас на глаза навертываются слезы в кино, театре, беседе?",
      "Сделав что-то, вы сомневаетесь, все ли сделано правильно и не успокаиваетесь до тех пор, пока не убедитесь еще раз в том, что все сделано правильно?",
      "В детстве вы были таким же смелым и отчаянным как ваши сверстники?",
      "Часто ли у вас меняется настроение от состояния безграничного ликования до отвращения к жизни, себе?",
      "Являетесь ли вы обычно центром внимания в обществе, компании?",
      "Бывает ли так, что вы беспричинно находитесь в таком ворчливом состоянии, что с вами лучше не разговаривать?",
      "Вы серьезный человек?",
	}
	
    myMapA := make(map[string][]int)
    myMapB := make(map[string][]int)
    myMapC := make(map[string][]int)
 
    myMapA["д"] = []int{1, 5}
    myMapA["н"] = []int{8}
 
    myMapB["д"] = []int{7}
    myMapB["н"] = []int{2, 4}
 
    myMapC["д"] = []int{3, 9}
    myMapC["н"] = []int{6}
    
    countA := 0
    countB := 0
    countC := 0
    
    questionCycle(&countA, &countB, &countC, myMapA, myMapB, myMapC, testQuestions)
    testResult(countA, countB, countC)

}
