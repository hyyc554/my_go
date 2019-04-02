package main

import "fmt"

func main() {
	// 冒泡排序
	a := [...]int{3, 5, 2, 1, 4, 7, 6}
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp

			}

		}
	}
	fmt.Println(a)
}
