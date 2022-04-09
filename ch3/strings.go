package main

import "fmt"

/**
notice that strings is immutable
**/
func main(){
	s := "left foot"
	t:= s
	s+= ", right foot"

	fmt.Println("s:",s)
	fmt.Println("t:", t)

}