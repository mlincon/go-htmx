/*
We are building a program that helps users track their expenses. We need to create
a struct called Expense to store information about an individual expense,
including the name of the expense, the amount, and the date.

We need to create a method called Total that calculates the total
amount spent on expenses.

Also, we need to create a method called getName on Expense struct that returns
the name of the Expense.
*/

package main

import "fmt"

// Expense struct to store information about an individual expense
type Expense struct {
	Name   string
	Amount float64
	Date   string
}

// Total method calculates the total amount spent on expenses
func Total(expenses []Expense) float64 {
	var total float64
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}

// getName method on the Expense struct returns the name of the Expense
func (e Expense) getName() string {
	return e.Name
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	// Print the total amount spent on expenses
	fmt.Println(Total(expenses))

	// Print the name of the first expense
	fmt.Println(expenses[0].getName())
}
