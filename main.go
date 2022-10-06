package main

import "fmt"

func LoopTopG(a string) string {
	for i := 0; i < 10000; i++ {
		fmt.Println("Andrew Tate is a Top G!")
	}
	return a
}

func main() {

	fmt.Println(LoopTopG("I told you so!"))

}
