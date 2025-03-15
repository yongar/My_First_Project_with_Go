package main

import "fmt"

func main() {
	// Write your code here
	bubblegum := 202
	bubblegum_Message := fmt.Sprintf("Bubblegum: $%d", bubblegum)
	toffee := 118
	toffee_Message := fmt.Sprintf("Ttoffee: $%d", toffee)
	ice_cream := 2250
	ice_cream_Message := fmt.Sprintf("Ice cream: $%d", ice_cream)
	milk_chocolate := 1680
	milk_chocolate_Message := fmt.Sprintf("Milk chocolate: $%d", milk_chocolate)
	doughnut := 1075
	doughnut_Message := fmt.Sprintf("Doughnut: $%d", doughnut)
	pancake := 80
	pancake_Message := fmt.Sprintf("Pancake: $%d", pancake)
	income := 0.0
	income = float64(bubblegum + toffee + ice_cream + milk_chocolate + doughnut + pancake)
	income_message := fmt.Sprintf("Income: $%.2f", income)
	fmt.Println("Earned amount:")
	fmt.Println(bubblegum_Message)
	fmt.Println(toffee_Message)
	fmt.Println(ice_cream_Message)
	fmt.Println(milk_chocolate_Message)
	fmt.Println(doughnut_Message)
	fmt.Println(pancake_Message)
	fmt.Println("")
	fmt.Println(income_message)
	fmt.Printf("%T", "1337")
	fmt.Println("")
	var staff_expense int
	var other_expense int
	fmt.Println("Staff expenses:")
	fmt.Scan(&staff_expense)
	fmt.Println("Other expenses:")
	fmt.Scan(&other_expense)
	net_income := income - float64(staff_expense) - float64(other_expense)
	fmt.Printf("Net income: $%.2f", net_income)
}
