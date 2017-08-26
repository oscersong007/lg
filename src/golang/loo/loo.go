/**
 * User: songjintao
 * Date: 2017/8/26
 * todo: 方法（面向对象）
 */

package loo

import "fmt"

func Run()  {
	//fmt.Print("loo")
	//var a N = 20
	//
	//a.value()
	//
	//a.pointer()
	//
	//fmt.Print("v:%p  %v",&a,a)

	var a N  = 25
	p:= &a

	a.value()
	a.pointer()

	p.value()
	p.pointer()

	println(a)
}

type N int

func (n N) value()  {
	n++
	fmt.Printf("v:%p  %v\n",&n,n)
}

func (n *N) pointer()  {
	(*n)++
	fmt.Printf("v:%p  %v\n",n,*n)
}
