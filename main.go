package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	choice := ""
	fmt.Println("Choose Calculation")
	fmt.Println("1 - Integration\n2 - Nl Curve Fitting\n3 - Root Finding: ")
	fmt.Scanln(&choice)

	if choice == "1" {
		initial := 0.0
		end := 0.0
		segment := 0

		fmt.Printf("Initial Point: ")
		fmt.Scanln(&initial)
		fmt.Printf("End Point: ")
		fmt.Scanln(&end)
		fmt.Printf("Number of Segments: ")
		fmt.Scanln(&segment)
		integration(initial, end, segment)
	} else if choice == "2" {
		fmt.Printf("X Array: ")
		strXArr := ""
		fmt.Scanln(&strXArr)
		strXArr = "[" + strXArr + "]"
		fmt.Printf("Y Array: ")
		strYArr := ""
		fmt.Scanln(&strYArr)
		strYArr = "[" + strYArr + "]"

		fmt.Println("Choose a Method")
		fmt.Println("1 - Power Model\n2 - Exponential Model\n3 - Saturation-Growth Model\n4 - Least Square Regression")
		method := 0
		fmt.Scanln(&method)

		var xArr []float64
		var yArr []float64
		err := json.Unmarshal([]byte(strXArr), &xArr)
		err = json.Unmarshal([]byte(strYArr), &yArr)
		if err != nil {
			log.Fatal(err)
		}
		calculation(xArr, yArr, method)
	} else if choice == "3" {
		fmt.Println("1 - Incremental Method\n2 - Bisection\n3 - False Position\n4 - Secant\n5 - Newton: ")
		fmt.Scanln(&choice)

		fmt.Printf("Number of iterations: ")
		iteration := 0
		fmt.Scanln(&iteration)
		fmt.Printf("Start At: ")
		start := 0.0
		fmt.Scanln(&start)
		if choice == "1" {
			fmt.Printf("Step size: ")
			step := 0.0
			fmt.Scanln(&step)
			incrementalSearch(iteration, start, step)
		} else if choice == "2" {
			fmt.Printf("End At: ")
			end := 0.0
			fmt.Scanln(&end)
			bisection(start, end, iteration)
		} else if choice == "3" {
			fmt.Printf("End At: ")
			end := 0.0
			fmt.Scanln(&end)
			falsePosition(start, end, iteration)
		} else if choice == "4" {
			fmt.Printf("End At: ")
			end := 0.0
			fmt.Scanln(&end)
			secant(start, end, iteration)
		} else if choice == "5" {
			newton(start, iteration)
		}

	} else {
		fmt.Println("Invalid Input!")
	}

}
