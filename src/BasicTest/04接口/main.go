package main

import "fmt"

type IMover interface {
	move()
	jump()
}

type dog struct {
	name string
}

func (d dog) move() {
	fmt.Println("狗会动")
}
func (d *dog) jump() {
	fmt.Println(d.name + " 在跳")
}
func main() {
	var x IMover
	var wangcai = dog{name: "旺财"} // 旺财是dog类型
	x = &wangcai                  // x可以接收dog类型
	x.move()
	var fugui = &dog{"富贵"} // 富贵是*dog类型
	x = fugui              // x可以接收*dog类型
	x.move()

	var y IMover
	var wangcai2 = &dog{name: "旺财02"} // 旺财是dog类型
	y = wangcai2                      // x不可以接收dog类型
	y.jump()
	var fugui2 = &dog{name: "富贵02"} // 富贵是*dog类型
	y = fugui2                      // x可以接收*dog类型
	y.jump()
}
