package main

import "fmt"

type Expense struct{
	name string
	price float64
	date string
}
func Total(expenses []Expense) float64{
	var total float64
	for _,expense:= range expenses{
		total+=expense.price
	}
	return total
}

func (expense *Expense) getName()string{
	return expense.name
}


func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}