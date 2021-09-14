package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动,Animal!\n", a.name)
}
func (a *Animal) jump() {
	fmt.Printf("%s会跳,Animal!\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s汪汪\n", d.name)
}
func (d *Dog) move() {
	fmt.Printf("%s会动,Dog!\n", d.Animal.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "大黄",
		},
	}
	d1.move()
	d1.Animal.move()
	d1.jump()
	d1.wang()
}
