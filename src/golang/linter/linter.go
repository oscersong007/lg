/**
 * User: songjintao
 * Date: 2017/8/26
 * todo: 接口(契约)
 */

package linter

import "fmt"

func Run()  {
	fmt.Print("interface...")

	var x C = "xx"
	var t N  = &x
	println(t)
}

type N interface {
	test()
	string() string
}

type C string

func (c C) test()  {
	fmt.Print("test")
}

func (c *C) string() string  {
	fmt.Print("string")
	return "xx"
}