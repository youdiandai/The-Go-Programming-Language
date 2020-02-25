/**
- 4.5：编写一个就地处理函数，用于去除 []string slice 中相邻的重复字符串元素
*/

package main

import "fmt"

func main() {
	a := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w", "a", "a"}
	a = removeMultiple(a)
	fmt.Println(a)
}

func removeMultiple(a []string) []string {
	//n := 0
	for i := 0; i < len(a); i++ {
		if a[i] == a[i+1] {
			a = append(a[:i], a[i+1:]...)
		}
	}
	return a
}
