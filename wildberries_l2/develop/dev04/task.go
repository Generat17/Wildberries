package main

import (
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func extractAnagrams(s []string) map[string][]string {
	groups := make(map[string][]string)

	for _, word := range s {
		word = strings.ToLower(word)
		key := sortString(word)
		groups[key] = append(groups[key], word)
	}

	res := make(map[string][]string, len(groups))
	for _, words := range groups {
		if len(words) > 1 {
			res[words[0]] = words
		}
	}
	return res
}

func sortString(s string) string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}
