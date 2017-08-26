/**
 * User: songjintao
 * Date: 2017/8/26
 * todo: 数据 array,slice,map
 */

package ldata

import (
	"fmt"
	"reflect"
)

func Run()  {
	fmt.Print("ldata...")
	//array()
	//slice()
	//dict()
	lstruc()
}

// 定长
// [len]type{init}
func array() {
	var a [4]int

	b:= [4]int{2,5}

	c:= [4]int{5,3:10}

	d:= [...]int{1,2,3,4}

	fmt.Print(a,b,c,d)
}

// 数组指针的封装
// []type{init}
func slice() {
	s := make([]int,0,5)
	for i:=1;i<200;i++ {
		s = append(s, i)
		// 不够时，分配当前的两倍
		fmt.Print(len(s),cap(s),"\n")
	}
}

func dict() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	println(m)
	if v,ok := m["b"];ok {
		println(v)
	}
	delete(m,"d")
}

func lstruc() {
	type user struct {
		name string `姓名`
		age  byte	`年龄`
	}

	u1 := user{name:"张三",age:12}
	u2 := user{name:"李四",age:34}
	fmt.Print(u1,u2)
	fmt.Print(reflect.ValueOf(u1).Type())
}