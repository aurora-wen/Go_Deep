package main

import (
	"fmt"
	"time"
)

func main() {
	values := make([]string, 6)
	values = append(values, "aa")
	values = append(values, "bb")
	values = append(values, "cc")
	values = append(values, "dd")
	err(values)
	//right(values)
}

/*
结论：
  - 最终协程会打印出values切片的最后一个值。

原因：
  - 因为当前val值引用的是同一个地址的数据，所以在range循环的过程中，会不断在val地址中更新数据；
  - 而在闭包中，由于引用了外部变量val，所以在访问时会获取val地址中的值，可能会获取最后放入其中的值，而不是遍历所有值，从而导致严重的错误。
*/
func err(values []string) {
	for _, val := range values {
		fmt.Println("in err func")
		go func() {
			fmt.Println("in err func closure func")
			fmt.Println(val)
		}()
		time.Sleep(2 * time.Second)
	}
}

/*
结论：
  - 通过函数传递参数，从而避免闭包引用导致的陷阱。
*/
func right(values []string) {
	for _, val := range values {
		go func(val string) {
			fmt.Println(val)
		}(val)
	}
}
