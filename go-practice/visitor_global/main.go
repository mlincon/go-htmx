/* We need to store the number of active visitors in a variable and update it
each time a new visitor arrives or an old visitor leaves the website.
*/

package main

import "fmt"

var activeUserCount int

func entry() {
	activeUserCount++
}

func exit() {
	activeUserCount--
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
