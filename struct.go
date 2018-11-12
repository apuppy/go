package main

import "fmt"

type person struct {
	name string
	age	 int
}

func main(){
	fmt.Println(person{"Kevin",26})

	fmt.Println(person{name: "Dewey",age: 23})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann",age: 30})

	s := person{name: "Sean", age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
