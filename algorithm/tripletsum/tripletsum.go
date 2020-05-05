package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//seed or get the same result
	//rand.Seed(time.Now().UnixNano())
	a := randomArray(2000)
	//fmt.Println(a)
	//rand.Shuffle(len(a), func(i, j int) {
	//	a[i], a[j] = a[j], a[i]
	//})
	//fmt.Println(a)
	//fmt.Println(TripletCount(a, 13))
	//fmt.Println(TripletCountEasy(a, 13))
	fmt.Println(timerWrapper(TripletCount, a, 13))
	fmt.Println(timerWrapper(TripletCountEasy, a, 13))
}

//用于测量函数运行时间的函数
func timer(start time.Time, id string) {
	elapsed := time.Since(start)
	fmt.Printf("%s time cost %s\n", id, elapsed)
}

type tripletFunc func([]int, int) int
//测量函数运行时间的另一种实现
func timerWrapper(f tripletFunc, arr []int, sum int) int {
	start := time.Now()
	result := f(arr, sum)
	elapsed := time.Since(start)
	fmt.Printf("time cost %s\n", elapsed)
	return result
}

func randomArray(size int) []int {
	s := make([]int, size, size)
	for i:=0; i<size;i++{
		s[i] = rand.Intn(10)
	}
	return s
}

//brute force solution
func TripletCount(arr []int, sum int) int {
	//defer timer(time.Now(), "n^3")
	l := len(arr)
	count := 0
	if l <= 2 {
		return count
	}
	for i:=0;i<l-2;i++{
		for j:=i+1;j<l-1;j++{
			for k:=j+1;k<l;k++{
				if arr[i]+arr[j]+arr[k]<sum{
					count++
				}
			}
		}
	}
	return count
}

func TripletCountEasy(arr []int, sum int) int {
	//defer timer(time.Now(), "n^2")
	l := len(arr)
	count := 0
	if l <= 2 {
		return count
	}
	sort.Ints(arr)
	for i:=0;i<l-2;i++{
		j := i+1
		k := l-1
		for j < k {
			if arr[i] + arr[j] + arr[k] >= sum{
				k--
			} else {
				count += k-j
				j++
			}
		}
	}
	return count
}
