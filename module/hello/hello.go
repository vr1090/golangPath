package main

import "example.com/greetings"
import "fmt"
import "log"



func main(){


	message, error := greetings.Hello("ujang")

	if error != nil {
		log.Fatal("error nih ", error)
	}

	fmt.Printf("%v", message)
}