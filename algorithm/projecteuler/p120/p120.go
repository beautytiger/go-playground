package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(p120(3, 1000))
	//fmt.Println(ReminderMax(700))
}

func p120(start, end int) int {
	var rMaxSum int
	for i:=start; i<=end; i++ {
		rMaxSum += ReminderMax(i)
	}
	return rMaxSum
}

func ReminderMax(in int) int {
	var a = big.NewInt(int64(in))
	var array []int64
	var power, aLeft, aRight, n big.Int
	var pointer, maxIndex int
	power.Exp(a, big.NewInt(2), nil)
	aLeft.Sub(a, big.NewInt(1))
	aRight.Add(a, big.NewInt(1))
	//rMax := 0
	n.SetInt64(1)
	for {
		ps := PowerSum(&aLeft, &aRight, &n)
		array = append(array, ps.Mod(ps, &power).Int64())
		if maxIndex !=0 && array[pointer] == array[maxIndex] {
			pointer++
		} else {
			pointer = 0
		}
		// 如果队列中有3个重复的数，则认为已经进入了循环
		if pointer >= 3{
			break
		}
		n.Add(&n, big.NewInt(1))
		maxIndex++
	}
	return int(ArrayMax(array))
}

//go居然没有好用的整数求幂运算符
func PowerSum(x, y, p *big.Int) *big.Int {
	var rx, ry big.Int
	rx.Exp(x, p, nil)
	ry.Exp(y, p, nil)
	return rx.Add(&rx, &ry)
}

//找到数组中最大的数
func ArrayMax(array []int64) int64 {
	var max int64 = 0
	for _, i := range array {
		if max < i {
			max = i
		}
	}
	return max
}
