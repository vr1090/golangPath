package main

import "fmt"

func main(){
	s := "left foot"
	t:= s
	s+= ", right foot"

	fmt.Println("s:",s)
	fmt.Println("t:", t)

}