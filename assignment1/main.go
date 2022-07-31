package main

import "fmt"

func main() {
	// //get the loop from 0 to 10
	// for i := 0; i <= 10; i++ {
	// 	//print in format
	// 	fmt.Printf("%d is ", i)
	// 	//find out if even or odd number
	// 	if i%2 == 0 {
	// 		fmt.Println("even")
	// 	} else {
	// 		fmt.Println("odd")
	// 	}
	// }

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is event")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
