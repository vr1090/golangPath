package main

import "fmt"

const(
	USD=iota
	EUR
	IDR
	INR
)

func main(){
	a:=[3]int{1,2,3}
	fmt.Println("a",a)

	var b [3]int
	fmt.Println("b",b)

	var c = [...]int{1,2,3,4,5,6}
	fmt.Println("c",c)

	curr := [...]int{INR:100,
		USD:99,
		EUR:78,
	}
	fmt.Println("curr",curr)
}