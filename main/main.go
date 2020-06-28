package main

import "fmt"

func main() {
	test := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", test))
}
