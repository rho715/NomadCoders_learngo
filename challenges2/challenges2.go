package challenges2

import (
	//"github.com/rho715/learngo/accounts"
	"fmt"

	"github.com/rho715/learngo/mydict"
)

// // 1. for accounts folder
// func main() {
// 	account := accounts.NewAccount("rhorho")
// 	account.Deposit(10)
// 	fmt.Println(account.Balance())
// 	err := account.Withdraw(25)
// 	if err != nil {
// 		fmt.Println(err)
// 		//log.Fatalln(err) // calls Println() & kills
// 	}
// 	account.ChangeOwner("euneun")
// 	//fmt.Println(account.Owner(), account.Balance())
// 	fmt.Print(account)
// }

// // 2. for mydict
// func main() {
// 	dictionary := mydict.Dictionary{"first": "First Word"}

// 	// search smt that is not inside dictionary
// 	definition, err := dictionary.Search("second")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(definition)

// 	// add a word
// 	word := "second"
// 	defin := "hehe"

// 	err2 := dictionary.Add(word, defin)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	// search a word
// 	secWord, _ := dictionary.Search(word)
// 	fmt.Println(secWord)

// 	// try to add again
// 	err3 := dictionary.Add(word, definition)
// 	if err3 != nil {
// 		fmt.Println(err3)
// 	}

// }

// 3. Delete
func challenges2() {
	// add a word to map
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	//update definition (hello:First -> hello:Second)
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	} // if word does not exists, print err msg
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word) // else print definition

	// update non exisiting word
	err2 := dictionary.Update("byoyoyoyd", "third")
	if err2 != nil {
		fmt.Println(err2)
	} // if word does not exists, print err2 msg
	word2, _ := dictionary.Search("byoyoyoyd")
	fmt.Println(word2) // else print definiiton

	// delete already existing word  & then search again
	dictionary.Delete(baseWord)
	wordResult, err3 := dictionary.Search(baseWord)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(wordResult)

}
