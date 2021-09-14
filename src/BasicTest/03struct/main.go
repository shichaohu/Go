package main

import "fmt"

type Persion struct {
	Name   string
	age    int8
	dreams []string
}

func (p *Persion) SetDreams(dreams []string) {
	//p.dreams = dreams
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func main() {
	p1 := Persion{Name: "小王", age: 18}
	data := []string{"吃饭", "sleep", "other"}
	p1.SetDreams(data)

	data[1] = "睡觉"
	fmt.Print((p1.dreams))

}
