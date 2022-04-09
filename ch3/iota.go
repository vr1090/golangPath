package main

import "fmt"

const (
	sunday = iota
	monday
	tuesday
)

func main(){
	fmt.Println(sunday, monday, tuesday)
}