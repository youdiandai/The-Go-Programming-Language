package main

import "fmt"

func main() {
	n := noZere(10)
	fmt.Println(n)

}

func noZere(n int) (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()
	panic(n)
}
