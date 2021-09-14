package main

import (
	"fmt"
	"sort"
)

func main() {
	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3] // s := a[low:high]
	// fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	// s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	// fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	//append
	// s1 := make([]int, 3, 5) //[0 0 0]
	// s2 := s1                //将s1直接赋值给s2，s1和s2共用一个底层数组
	// s2[0] = 100
	// s2 = append(s2, 7, 8)

	// s3 := append(s1, s2...) // [1 2 3 4 5 6 7]

	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println(s3)

	//append
	fmt.Println("append-----------------------------")
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	newSlice = append(newSlice, 70, 80, 90, 100)
	fmt.Println(slice, newSlice)

	//排序
	fmt.Println("排序-----------------------------")
	var a = [...]int{3, 7, 8, 9, 1}
	// fmt.Println(a[:]) //将数组切片
	//sort包内部实现了内部数据类型的排序
	//升序排列
	sort.Ints(a[:])
	fmt.Println(a)
	//降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)

	//map:
	fmt.Println("map-----------------------------")
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["小龙"] = 110
	//等效于：
	// scoreMap := make(map[string]int){
	// 	"张三":  90
	// 	"小明":  100
	// }

	// 如果scoreMap的key中存在，则ok为true,v为对应的值；不存在，则ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	fmt.Print("v%\n", v)
	fmt.Print("v%\n", ok)
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	//遍历map时的元素顺序与添加键值对的顺序无关。
	for k, v := range scoreMap {
		fmt.Print(k, v)
	}
	fmt.Println("删除小明")
	delete(scoreMap, "小明") //将小明:100从map中删除
	//只遍历key
	for k := range scoreMap {
		fmt.Print(k)
	}
}
