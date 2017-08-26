/**
 * User: songjintao
 * Date: 2017/8/26
 */

package lcon

import (
	"fmt"
	"sync"
)

func Run() {
	s1()
}

// select
func s1()  {
	var wg sync.WaitGroup
	wg.Add(3)

	a,b := make(chan int),make(chan int)

	go func() {
		defer wg.Done()
		for  {
			select {
			case x,ok:=<-a:
				if !ok {
					a = nil
					break
				}
				fmt.Print("a:", x,"\n")
			case x,ok:=<-b:
				if !ok {
					b = nil
					break
				}
				fmt.Print("b:",x,"\n")
			}
			if a == nil && b== nil {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		for i:=7;i<10 ;i++  {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)
		for i:=0;i<3 ;i++  {
			b <- i
		}
	}()

	wg.Wait()

	fmt.Print("exit")
}

// channel
func c1() {
	done := make(chan int)
	c := make(chan int)

	go func() {
		defer close(done)

		for x:= range c{
			fmt.Print(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}
