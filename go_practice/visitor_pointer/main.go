/* We need to store the number of active visitors in a variable and update it
each time a new visitor arrives or an old visitor leaves the website.
*/

package main

import "fmt"

// Declare variable activeUserCount as a pointer to int
var activeUserCount *int

func entry() {
	// Increment the value pointed by activeUserCount when a new visitor arrives
	*activeUserCount++
}

func exit() {
	// Decrement the value pointed by activeUserCount when a visitor leaves
	*activeUserCount--
}

func main() {
	// Initialize activeUserCount with a new integer variable
	activeUserCount = new(int)

	entry()
	entry()
	exit()
	exit()
	entry()
	entry()

	// Dereference the pointer to get the actual value of activeUserCount
	fmt.Println(*activeUserCount)
}
