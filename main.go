package main

import (
	"fmt"

	"github.com/JaesunJin/mydict"
)

func main() {
	dict := mydict.Dictionary{}

	word := "Hello"
	def := "Greeting"

	err := dict.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dict.Search(word))

	if dict.Update(word, "Bye") != nil {
		fmt.Println(err)
	}

	fmt.Println(dict.Search(word))

	err = dict.Update("test", "Bye")
	if err != nil {
		fmt.Println(err)
	}

	err = dict.Delete("test")
	if err != nil {
		fmt.Println(err)
	}

	err = dict.Delete(word)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dict.Search(word))
}
