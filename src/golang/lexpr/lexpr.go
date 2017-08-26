/**
 * User: songjintao
 * Date: 2017/8/25
 * todo: 表达式
 */

package lexpr

import (
	"fmt"
)

func Run()  {

	fmt.Print("lexpr...")

	// 25个保留关键字

	// 指针
	x := 1;
	var p *int  = &x
	*p += 20
	fmt.Print(p,*p)

	// 初始化
	// 1. 必须含类型标签
	// 2. 左边{不能另起一行
	// 3. 多个初始化,分割
	// 4. 允许多行但是必须以,或者}结束
	type data struct {
		x int
		s string
	}

	var a data = data{1,"123"}
	fmt.Print(a)

	// 流控制
	// if
	if x < 3 {
		fmt.Print("small 3")
	} else if x < 50 && x > 10 {
		fmt.Print("<50 and >10")
	} else {
		fmt.Print("if_over.")
	}

	// switch
	switch x {
	case 1:
		fmt.Print("a")
	case 21:
		fmt.Print("21")
		fallthrough
	default:
		fmt.Print("default")
	}

	// for
	for i:=0; i<10; i++ {
		if i > 8 {
			break
		}
		if i % 4 == 0 {
			continue
		}
		fmt.Print(i)
	}

	ss :=[3]string{"a","b","c"}
	for i,s :=range ss{
		fmt.Print(i,s)
	}


}
