package main

import "fmt"

type Employee struct{
	name string
	id int
	position string
}

func (e Employee) string() string{
	return fmt.Sprint("%s :: %s :: %d ", e.name, e.position,
	 e.id)
}

func main(){
	var k = Employee{name:"inu", 
		id:1,position:"entry"}
	fmt.Println(k)
	var pos = &k.position
	*pos = "senior CEO"
	fmt.Println(k)
	var j = &k
	fmt.Println("with a pointer ",j)
	fmt.Println("try to change the value")
	j.name = "ubahan inu"
	fmt.Println("habis di ubah", j)
}