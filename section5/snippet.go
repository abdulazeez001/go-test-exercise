package main

import (
	"fmt"
	"math/big"
)

// ISSUES
// 1. factorial of a number is the product of all positive integers up to the number.
//    but it currently performs addition instead of multiplication  at [result += 1].
//    though it's possible to get the factorial of decimal(float). this function formula doesn't care for that.

// 2. At [result := 1.00] The result variable is initialized as a float64 instead of an int.
//    Factorials are usually integers.
//    though it's possible to get the factorial of decimal(float). this function formula doesn't care for that.

// 3. The printfactorial() is not called properly it should be printFactorial()

// 4. The value of factorial for a large number can be to big for an integer, though it is
//    okay for small numbers

// 5. printFactorial() Should accept num has a parameter so that it can be reusable for other
//    numbers

// RECOMENDATIONS
// 1. Use multiplication instead of addition
// 2. Use integer instead of float64
// 3. The printfactorial() is not called properly it should be printFactorial()
// 4. We can use big.Int as the number type to prepare for very large number
// 5. printFactorial() Should accept num has a parameter so that it can be reusable for other
//    numbers
// 6. We can also make use of goroutine to make the calculation faster

// ******** ERROR SOLUTION ************ //

// Function to calculate the factorial of a number
// func factorial(n int) int {
// 	result := 1.00
// 	for i := 1; i <= n; i++ {
// 		result += i
// 	}
// 	return result
// }

// // Function to print the factorial of a number
// func printFactorial() {
// 	num := 5
// 	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
// }

// ************  RIGHT SOLUTION *********** //
//
// func factorial(n int) int {
// 	result := 1 // Initialize result as an integer (was previously float64, which is unnecessary for factorials)
// 	for i := 1; i <= n; i++ {
// 			result *= i // Corrected the operation to multiplication (previously addition)
// 	}
// 	return result
// }

// // Function to print the factorial of a number
// func printFactorial(num int) {
// 	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num)) // Using the corrected factorial function
// }

// ************  BIG INT SOLUTION *********** //

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for index := int64(1); index <= n; index += 1 {
		result.Mul(result, big.NewInt(index))
	}
	return result
}

// Function to print the factorial of a number
func printFactorial(num int64) {
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}

func main() {
	printFactorial(2)
	printFactorial(60)
}

// func main() {
// 	printfactorial()
// }
