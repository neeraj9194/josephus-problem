package main

import (
	"fmt"
)

func main() {
	var numPerson, skipStep int
	fmt.Printf("Enter the number of persons in circle: ")
	_, err := fmt.Scan(&numPerson)
	if err != nil {
		fmt.Println("Please enter a valid integer")
		return
	}

	fmt.Printf("Enter the number of skip before kill (kth value-1): ")
	_, err = fmt.Scan(&skipStep)
	if err != nil {
		fmt.Printf("Please enter a valid integer :%v:", skipStep)
		return
	}

	circle := new(Circle)
	circle.Init(numPerson, skipStep)
	alivePerson := circle.SimulateKills()
	fmt.Printf("Person finally alive: %v\n", alivePerson.pos+1)

}
