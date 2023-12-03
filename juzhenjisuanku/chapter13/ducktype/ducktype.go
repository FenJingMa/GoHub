package ducktype

import "fmt"

//定义一个接口interface，抽象出鸭子的三个重要特征：叫喊，行走，游泳
type DuckType interface {
	Shout()
	Walk(road string)
	Swim(river string)
}

//利用结构体定义鸭子的属性
type Duck struct {
	Name   string
	Weight float64
}

//利用方法来定义鸭子的三种行为
//注意Go语言中的方法，类型变量写在func关键字之前。类型变量使用字母简写
func (d Duck) Shout() {
	fmt.Println("my name is:", d.Name, " ga ga ga")
}

func (d Duck) Walk(road string) {
	fmt.Println(d.Name, "is walking on the", road, " ga ga ga")
}

func (d Duck) Swim(river string) {
	fmt.Println(d.Name, "is swimming in the", river, " ga ga ga")
}

//利用结构体来定义人的基本属性
type Person struct {
	Name   string
	Weight float64
	Height float64
}

//利用方法来定义人，注意到人的三个方法具体实现的输出结果，与鸭子明显不同
func (p Person) Shout() {
	fmt.Println("this person name is:", p.Name)
}

func (p Person) Walk(road string) {
	fmt.Println(p.Name, "is walking on the", road)
}

func (p Person) Swim(river string) {
	fmt.Println(p.Name, "is swimming in the", river)
}
