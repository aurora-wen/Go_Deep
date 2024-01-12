package main

import "fmt"

func main() {
	/*
		单精读浮点数float32的精读为6~8位；
		双精度浮点数float64的精读为15~17位；
	*/
	noMissingFloat32 := float32(0.3)
	missingFloat32 := float32(0.333333336)
	noMissingFloat64 := 0.33333333333333336
	missingFloat64 := 0.3333333333333333336
	fmt.Println("noMissingFloat32: ", noMissingFloat32)
	fmt.Println("missingFloat32: ", missingFloat32)
	fmt.Println("--------")
	fmt.Println("noMissingFloat64: ", noMissingFloat64)
	fmt.Println("missingFloat64: ", missingFloat64)

	fmt.Println("=======================================")
	/*
		加法运算由于需要进行指数调整，有丢失精度的风险；
		在涉及 加、减、乘和除的运算时：
			* 优先执行乘法和除法运算；
			* a x (b+c) 可以转化为 axb + axc
	*/
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println(n)
}
