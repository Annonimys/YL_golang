package main

import (
	"fmt"
	"sort"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range wordMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var ans []string
	for i := 0; i < n; i++ {
		ans = append(ans, ss[i].Key)
	}
	return ans
}

func AnalyzeText(text string) {
	r := strings.NewReplacer(
		",", " ",
		".", " ",
		"!", " ",
		"?", " ",
	)
	text = r.Replace(text)
	words := strings.Split(text, " ")

	wordMap := make(map[string]int)
	cnt := 0
	for _, word := range words {
		if word == "" {
			cnt++
			continue
		}

		wordMap[strings.ToLower(word)]++
	}

	ans := getTopWords(wordMap, 5)

	fmt.Printf("Количество слов: %d\n", len(words)-cnt)
	fmt.Printf("Количество уникальных слов: %d\n", len(wordMap))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", ans[0], wordMap[ans[0]])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for i := 0; i < 5; i++ {
		fmt.Printf("\"%s\": %d раз\n", ans[i], wordMap[ans[i]])
	}
}
