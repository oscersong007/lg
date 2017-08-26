/**
 * User: songjintao
 * Date: 2017/8/25
 * @todo 类型
 */

package ltype

import (
	"fmt"
	"strconv"
)

func Run()  {

	// 类型在变量之后
	var xx int
	xx = 123
	fmt.Print(xx)

	// 简短模式
	y := "short define.\n"
	fmt.Print(y)

	// 多变量复制
	a,b := 1,2
	a,b = a+1,b+1
	fmt.Print(a,b)

	// 命名：驼峰.

	// 空标识符：_
	c,_ := strconv.Atoi("1234")
	fmt.Print(c)

	// 常量
	const (
		aa int = 1
		bb
		cc string = "c"
		dd
	)
	fmt.Print(aa,bb,cc,dd)

	// 枚举
	const (
		_ = iota
		kb = 1 << (10*iota)
		mb
		gb
	)
	fmt.Print(kb,mb,gb)

	// 别名：
	// byte ==> uint8
	// rune ==> int32

	// 引用类型： slice, map, channel
	// new : 分配零值内存，返回指针
	// make : 相关属性初始化

	// 强制转化
	type t int
	var tVal t
	tVal = 1
	fmt.Print(tVal)
}
