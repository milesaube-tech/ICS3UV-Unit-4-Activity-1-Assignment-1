/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks the user for a number of marks, reads each mark using a loop,
 */

package main

import (
	"fmt"
)

func main() {
	// get user input
	var numMarks int
	fmt.Print("How many marks will you enter today? ")
	fmt.Scan(&numMarks)

	// set variables
	var total float64 = 0
	var mark float64 = 0

	// read in each mark
	for i := 1; i <= numMarks; i++ {
		fmt.Printf("Enter mark %d: ", i)
		fmt.Scan(&mark)
		total = total + mark
	}

	// calculate average
	average := total / float64(numMarks)

	// display results
	fmt.Printf("You have entered %d marks. The student's average is %.1f%%.\n", numMarks, average)

	// give feedback
	if average <= 49 {
		fmt.Println("The student is failing.")
	} else if average >= 50 && average <= 69 {
		fmt.Println("The student's performance is just under average.")
	} else if average >= 70 && average <= 79 {
		fmt.Println("The student's performance is average.")
	} else if average >= 80 {
		fmt.Println("The student is on the honour roll.")
	}
}