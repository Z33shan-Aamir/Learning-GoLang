package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add(values []string) float32 {

	var sum float32 = 0
	for i := 0; i < len(values); i++ {
		value := values[i]
		value_float, err := strconv.ParseFloat(value, 64)
		if err == nil {
			sum += float32(value_float)
		} else {
			fmt.Println("Yeah `Add()` fucked up during the conversion")
		}

	}
	return sum
}

func Subtract(values []string) float32 {
	if len(values) > 2 && len(values) < 2 {
		fmt.Println("Please Enter only two values")
		return 0
	} else {
		var subtract float32 = 0
		for i := 0; i < len(values); i++ {
			value := values[i]
			value_float, err := strconv.ParseFloat(value, 64)
			if err == nil {
				if i == 0 {
					subtract = float32(value_float)
				} else {
					subtract -= float32(value_float)

				}
			} else {
				fmt.Println("Yeah I fucked up during the conversion")
			}

		}

		return subtract
	}

}

func Multiply(values []string) float32 {
	if len(values) < 2 {
		fmt.Println("Please enter at least two values to multiply")
		return 0
	}
	var product float32 = 1
	for i := 0; i < len(values); i++ {
		value := values[i]
		value_float, err := strconv.ParseFloat(value, 64)
		if err == nil {
			product *= float32(value_float)
		} else {
			fmt.Println("Error during conversion in Multiply:", err)
		}
	}
	return product
}

func Divide(values []string) float32 {
	if len(values) != 2 {
		fmt.Println("Please enter exactly two values to divide")
		return 0
	}
	var result float32
	for i := 0; i < len(values); i++ {
		value := values[i]
		value_float, err := strconv.ParseFloat(value, 64)
		if err == nil {
			if i == 0 {
				result = float32(value_float)
			} else {
				if value_float == 0 {
					fmt.Println("Error: Division by zero")
					return 0
				}
				result /= float32(value_float)
			}
		} else {
			fmt.Println("Error during conversion in Divide:", err)
		}
	}
	return result
}
func main() {

	fmt.Println("My first go App")
	fmt.Println("Do you want to 'add', 'subtract', 'multiply', or 'divide' value")
	var option string
	fmt.Scanln(&option)
	option = strings.ToLower(option)
	switch option {
	case "add":
		fmt.Println("Enter all the values you want to add seperated by white space")
	case "subtract":
		fmt.Println("Enter 1st and 2nd value to subtract seperated by white space")
	case "multiply":
		fmt.Println("Enter 1st and 2nd value to multiply seperated by white space")
	case "divide":
		fmt.Println("Enter 1st and 2nd value to divide seperated by white space")
	}

	// Use bufio.NewReader to read the entire line of input
	reader := bufio.NewReader(os.Stdin)
	user_input, _ := reader.ReadString('\n')   // Read until a newline character
	user_input = strings.TrimSpace(user_input) // Remove any trailing newline or spaces

	// fmt.Println("User Input:", user_input)

	// Split the input into values
	values := strings.Fields(user_input)
	// fmt.Println("Values:", values)
	// var results float32 = Add(values)
	// fmt.Println("Sum is:", results)

	switch option {
	case "add":
		fmt.Println(Add(values))
	case "subtract":
		fmt.Println(Subtract(values))
	case "multiply":
		fmt.Println(Multiply(values))
	case "divide":
		fmt.Println(Divide(values))
	}

}
