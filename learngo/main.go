package main

import (
	"fmt"

	"github.com/icefirebear/learngo/mydict"
)

func main() {
	//Account Example Code
	// account := accounts.NewAccount("woung")
	// account.Deposit(10)
	// fmt.Println(account.Balance())
	// // err := account.Withdraw(20)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// fmt.Println(account)

	//Dictionary Example Code
	// dictionary := mydict.Dictionary{"first": "First Word"}
	// word := "hello"
	// definition := "greeting"

	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	//Update method Example Code
	// dictonary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictonary.Add(baseWord, "Second")
	// err := dictonary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictonary.Search(baseWord)
	// fmt.Println(word)

	// Delete Method Example Code
	dictonary := mydict.Dictionary{}
	baseWord := "hello"
	dictonary.Add(baseWord, "Second")
	dictonary.Search(baseWord)
	dictonary.Delete(baseWord)
	word, err := dictonary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
