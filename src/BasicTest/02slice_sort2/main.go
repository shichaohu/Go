package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	persons := []Person{
		{"zhang3", 10},
		{"li4", 50},
		{"wang5", 30},
	}

	//Go 1.8中，加入了sort.Slice函数
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	})
	fmt.Print("v%\n", persons)
}
