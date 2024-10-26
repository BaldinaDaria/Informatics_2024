package main

import (
	"fmt"
)

type Cat struct {
	name  string 
	age   int    
	breed string 
}

func (c *Cat) GetAge() int {
	return c.age
}

func (c *Cat) SetAge(age int) {
	c.age = age
}

func (c *Cat) Info() string {
	return fmt.Sprintf("Кошка: %s, Возраст: %d, Порода: %s", c.name, c.age, c.breed)
}

func main() {
	
	myCat := Cat{name: "Барсик", age: 7, breed: "Мейн-кун"}

	fmt.Println("Возраст кошки:", myCat.GetAge())

	myCat.SetAge(8)
	fmt.Println("Новый возраст кошки:", myCat.GetAge())

	fmt.Println(myCat.Info())
}
