package examples

import "fmt"

// Loop loops over data in different ways.
func Loop() {
	// Looping - while loop
	count := 0
	for {
		if count == 10 {
			fmt.Print("Exit")
			break
		}
		count++
	}

	// Looping - for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}

	arr := []int{10, 20, 30, 40, 50, 60}
	arr = append(arr, 70)

	// Looping - arrays
	for index, val := range arr {
		fmt.Printf("%v\n%v\n", index, val)
	}
}
