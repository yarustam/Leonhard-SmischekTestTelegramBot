package main

import (
	"fmt"
	"sort"
)

func testResult(counts []int) {
	groups := []string{
		"Демонстративность",
		"Застревание",
		"Педантичность",
		"Возбудимость",
		"Гипертимность",
		"Дистимность",
		"Тревожность",
		"Экзальтированность",
		"Эмотивность",
		"Циклотимность"}

	multipliers := []int{2, 2, 2, 3, 3, 3, 3, 6, 3, 3}
	result := make(map[string]int)

	for i, count := range counts {
		result[groups[i]] = count * multipliers[i]
	}

	//сортировка результатов в порядке убывания

	//сперва создадим структуру
	type kv struct {
		Key   string
		Value int
	}

	var sorted []kv
	for k, v := range result {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	for _, kv := range sorted {
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
