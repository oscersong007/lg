/**
 * User: songjintao
 * Date: 2017/8/26
 * todo: 函数
 */

package lfunc

import (
	"fmt"
	"errors"
)

func Run()  {
	fmt.Print("lfuncxx")
	f:=hello
	exec(f)

	fmt.Print(add(1,2))

	a := []int{1,2,3}
	test(a...)

	add := rfunc()
	fmt.Print(add(12,34))

	fmt.Print("\n")

	fmt.Print(rdefer())

	rdefer2()

	if _,err:=div(1,0);err!=nil {
		println(err)
	}

	pfunc()
}

// 支持定长变参
// 多返回值
// 不支持默认传参

func hello() {
	fmt.Print("hello world")
}

func exec( f func())  {
	f()
}

func add(a,b int)  (int,int)  {
	return a+b,a*b
}

// 变参
func test(a ...int)  {
	for _,i := range a {
		fmt.Print(i)
	}
}


// 函数作为返回值
func rfunc() func(a,b int) int {
	return func(a, b int) int {
		return a+b
	}
}

// 延迟函数 (filo) 可修改return值
func rdefer() (z int)  {
	defer func() {
		println("defer:",z)
		z+=100
	}()
	return 100
}

func rdefer2()  {
	x,y := 1,2
	defer func(a int) {
		println("defer",a,y)
	}(x)	// 注册时复制参数
	x += 100
	y += 200
	fmt.Print(x,y)
}

// 错误处理
func div(x,y int)(int,error)  {
	if y == 0 {
		return 0,errors.New("div by zero")
	}
	return x/y,nil
}

// panic recover
func pfunc() {
	defer func() {
		if err:=recover();err!=nil {
			println("recover")
			println(err)
		}
		println("ok")
	}()
	// panic("panic")
	println("exit.")
}
