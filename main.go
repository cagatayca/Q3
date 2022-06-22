package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func word_count(words []string) map[string]int {

	word_freq := make(map[string]int)

	for _, word := range words {
		_, ok := word_freq[word]
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}

type word_struct struct {
	freq int
	word string
}

func (p word_struct) String() string {
	return fmt.Sprintf(p.word)
}

type by_freq []word_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq < a[j].freq }

func main() {

	text := `"apple", "pie", "apple", "red", "red", "red""`

	text = strings.ToLower(text)

	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(text, -1)

	word_map := word_count(words)

	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	ix := 0
	for key, val := range word_map {
		pws.freq = val
		pws.word = key
		struct_slice[ix] = *pws
		ix++
	}

	sort.Sort(by_freq(struct_slice))

	fmt.Println(struct_slice[len(struct_slice)-1])

}
