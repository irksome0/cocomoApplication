package main

import (
	"fmt"
	"math"
	"strconv"
)

var TYPE_ORGANIC = [...]float32{2.4, 1.05, 2.5, 0.38}
var TYPE_SEMIDETACH = [...]float32{3.0, 1.12, 2.5, 0.35}
var TYPE_EMBEDDED = [...]float32{3.6, 1.20, 2.5, 0.32}

func main() {
	var modelType string
	for {
		fmt.Print("Enter a model type(basic/intermediate/advanced)\n B/I/A : ")
		fmt.Scan(&modelType)
		switch modelType {
		case "B":
			basicModel()
			return
		case "I":
			return
		case "A":
			return
		default:
			fmt.Println("Try again...")
		}
	}
}

func basicModel() {
	var projectType string
	var coefficients [4]float32
	for {
		fmt.Print("Enter a project type(organic/semidetach/embedded)\n O/S/E : ")
		fmt.Scan(&projectType)
		flag := false
		switch projectType {

		case "O":
			coefficients = TYPE_ORGANIC
			flag = true
			break
		case "S":
			coefficients = TYPE_SEMIDETACH
			flag = true
			break
		case "E":
			coefficients = TYPE_EMBEDDED
			flag = true
			break
		default:
			fmt.Println("Try again...")
		}
		if flag {
			break
		}
	}
	var SIZE string
	fmt.Print("What's the size of program: ")
	fmt.Scan(&SIZE)
	size, err := strconv.ParseFloat(SIZE, 64)
	if err != nil {
		panic(err)
	}
	PM := float64(coefficients[0]) * math.Pow(size, float64(coefficients[1]))
	TM := float64(coefficients[2]) * math.Pow(PM, float64(coefficients[3]))
	SS := PM / TM
	P := size / float64(PM)
	fmt.Println(PM, TM, SS, P, coefficients)
}
