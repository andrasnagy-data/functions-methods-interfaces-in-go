/*
Displacement = 1/2 * a * t**2 + v0 * t + s0

a = acceleration
t = time
v0 = initial velocity
s0 = initial displacement

*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var values string
	// ask user fo a, v0, and s0
	fmt.Println(
		"Enter 3 comma separated integer values for a, v0, and s0 (e.g. 10,2,1): ")
	fmt.Scan(&values)

	// Parse values into a flaot64 slice
	numbersStrSlc := strings.Split(values, ",")
	var numbersSlc []float64
	for _, elem := range numbersStrSlc {
		digit, _ := strconv.Atoi(elem)
		numbersSlc = append(numbersSlc, float64(digit))
	}

	// ask user for t
	var timeStr string
	fmt.Println("Enter an integer values for t: ")
	fmt.Scan(&timeStr)

	// parse timeStr into a float64
	num, _ := strconv.Atoi(timeStr)
	t := float64(num)

	// get a parametrized function
	fn := GenDisplaceFn(numbersSlc[0], numbersSlc[1], numbersSlc[2])

	// compute and print displacement
	fmt.Println(fn(t))

}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}
