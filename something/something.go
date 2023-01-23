package something

import "fmt"

func sayBye() { //private function
	fmt.Println("Bye") 
}

func SayHello() { //exported function (start with UpperCase)
	fmt.Println("Hello")
}