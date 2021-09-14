package main

import (
	"fmt"
	"sort"
)

/*slice 排序示例*/
type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (s PersonSlice) Len() int           { return len(s) }
func (s PersonSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	persons := PersonSlice{
		Person{
			Name: "甲",
			Age:  1,
		},
		Person{
			Name: "乙",
			Age:  5,
		},
		Person{
			Name: "丙",
			Age:  2,
		},
	}
	sort.Sort(persons)
	fmt.Printf("after sort:%+v", persons)

	//对象未数组时，以上代码等效于 sort.Ints ,例如：
	s := []int{89, 14, 8, 9, 17, 56, 95, 3}
	sort.Ints(s)
}
