package main

import (
	"fmt"
	"math/big"
)

/*
	math/big标准库提供了处理大数的三种数据类型：
		* big.Int
		* big.Float
		* big.Rat

	-------------------Int--------------------
	其中Int结构：
	type Int struct {
		neg bool  //符号
		abs nat   //整数位
	}
	type nat []Word
	type Word unit

	big.Int的核心思想：用uint切片来存储大整数，可以容纳的数据超过了int64的大小，甚至可以认为它是可以无限扩容的。
	-------------------Int--------------------

	-------------------Float--------------------
	其中Float结构：
	type Float struct {
		prec uint32  		//存储数字的位数
		mode RoundingMode
		acc  Accurarcy
		form form
		neg  bool			//符号位
		mant nat			//大整数
		exp  int32			//指数
	}

	big.Float的核心思想：把浮点数转换为大整数运算。
	-------------------Float--------------------

	-------------------Rat--------------------
	其中Rat结构：
	type Rat struct {
		a, b big.Int
	}
	big.Rat的核心思想：将有理数的运算转换为分数的运算。
	-------------------Rat--------------------
*/

func main() {
	bigInt()
	bigFloat()
	bigRat()
}

func bigInt() {
	// big.NewInt函数初始化big.Int
	a := big.NewInt(0)
	b := big.NewInt(1)
	var limit big.Int
	// 使用big.Exp函数计算 10^99 的大小
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)
	// 使用big.Cmp函数比较大整数的值
	for a.Cmp(&limit) < 0 {
		// 使用big.Add函数计算大整数的加法
		a.Add(a, b)
		a, b = b, a
	}
	fmt.Println(a)
	fmt.Println("==========================")
}

func bigFloat() {
	var x1, y1 float64 = 10, 3
	z1 := x1 / y1
	fmt.Println(x1, y1, z1)
	x2, y2 := big.NewFloat(10), big.NewFloat(3)
	z2 := new(big.Float).Quo(x2, y2)
	// 不设置prec时，其精度与float64相同
	fmt.Println(x2, y2, z2)

	// prec越大，精度越准，消耗也越大
	// 设置prec后，精度提升
	x2.SetPrec(100)
	y2.SetPrec(100)
	fmt.Println(x2, y2, z2)
	fmt.Println("==========================")
}

func bigRat() {
	// 如果希望有理数计算不丢失精度，那么可以借助big.NewRat实现
	x, y := big.NewRat(1, 2), big.NewRat(2, 3)
	z := new(big.Rat).Mul(x, y)
	fmt.Println("z: ", z)
}
