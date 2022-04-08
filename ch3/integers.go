package main

import (
	"fmt"
)

func main(){
	var apple int32 = 1
	var orange int64 = 1
	total := int(apple) + int(orange)

	fmt.Printf("total bang %d", total)
}