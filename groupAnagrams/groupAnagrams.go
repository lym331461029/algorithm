package main

import (
	"fmt"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//strs := []string{""}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

func toSortedAnagrams(str string) string {
	a := NewAnagram(str)
	return sortedAnagram(a)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		sortedAnagramStr := toSortedAnagrams(str)
		m[sortedAnagramStr] = append(m[sortedAnagramStr], str)
	}

	result := make([][]string, 0, len(m))
	for _, strg := range m {
		result = append(result, strg)
	}

	return result
}

type anagram struct {
	alphabets [26]int8
	rawStr    string
}

func sortedAnagram(a *anagram) string {
	//sb := strings.Builder{}
	//for i, c := range a.alphabets {
	//	for ; c > 0; c-- {
	//		sb.WriteByte(byte(i + 'a'))
	//	}
	//}
	//return sb.String()

	rawStr := make([]byte, 0, len(a.rawStr))
	for i, c := range a.alphabets {
		for ; c > 0; c-- {
			rawStr = append(rawStr, byte(i+97))
		}
	}
	return string(rawStr)
}

func NewAnagram(str string) *anagram {
	a := &anagram{
		rawStr: str,
	}
	for _, alpha := range str {
		a.alphabets[int8(alpha)-97] += 1
	}
	return a
}
