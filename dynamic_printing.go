package main

import "fmt"

func main() {
	var i int = 1
	var t bool = true
	var a int = 1
	for t {
		i = i + 1
		fmt.Print("\rHello World : ", i)
		if i == 1000000 || i%1000000 == 0 {
			fmt.Println("\ndo you wanna stop this(1 for yes 0 for No): ")

			fmt.Scanln(&a)
			if a == 1 {
				t = false
			} else {
				t = true
			}
		}

	}
}
