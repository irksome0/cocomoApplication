package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/irksome0/cocomoApplication/data"
)

func main() {
	var modelType string
	for {
		fmt.Print("Enter a model type(basic/intermediate)\n B/I : ")
		fmt.Scan(&modelType)
		switch modelType {
		case "B":
			basicModel()
			return
		case "I":
			intermediateModel()
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
			coefficients = data.BASIC["ORGANIC"]
			flag = true
			break
		case "S":
			coefficients = data.BASIC["SEMIDETACH"]
			flag = true
			break
		case "E":
			coefficients = data.BASIC["EMBEDDED"]
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
	fmt.Printf("Labor intensity: %f\nDevelopment duration: %f\nAmount of developers: %f\nEfficiency: %f\n", PM, TM, SS, P)
}
func intermediateModel() {

	var CD = make([]string, 15, 15)
	var projectType string
	var coefficients [2]float32
	var SIZE string

	for {
		fmt.Print("Enter a project type(organic/semidetach/embedded)\n O/S/E : ")
		fmt.Scan(&projectType)
		flag := false
		switch projectType {

		case "O":
			coefficients = data.INTERMEDIATE["ORGANIC"]
			flag = true
			break
		case "S":
			coefficients = data.INTERMEDIATE["SEMIDETACH"]
			flag = true
			break
		case "E":
			coefficients = data.INTERMEDIATE["EMBEDDED"]
			flag = true
			break
		default:
			fmt.Println("Try again...")
		}
		if flag {
			break
		}
	}

	fmt.Print("What's the size of program: ")
	fmt.Scan(&SIZE)

	size, err := strconv.ParseFloat(SIZE, 64)
	if err != nil {
		panic(err)
	}
	fmt.Print("Enter coefficients(1-6): \n")
	for i := 0; i < 15; i++ {
		fmt.Print(data.ATTRIBS[i] + ": ")
		var temp string
		fmt.Scan(&temp)
		switch temp {
		case "1":
			CD[i] = "VERY LOW"
			break
		case "2":
			CD[i] = "LOW"
			break
		case "3":
			CD[i] = "INTERMEDIATE"
			break
		case "4":
			CD[i] = "HIGH"
			break
		case "5":
			CD[i] = "VERY HIGH"
			break
		case "6":
			CD[i] = "CRITICAL"
			break
		default:
			CD[i] = "INTERMEDIATE"
			break
		}
	}

	productOfCoefs := 1.0
	for i := 0; i < 15; i++ {
		productOfCoefs *= data.ATTRIBUTES[data.ATTRIBUTES_TAGS[i]][CD[i]]
	}

	PM := productOfCoefs * float64(coefficients[0]) * math.Pow(size, float64(coefficients[1]))
	fmt.Printf("Labor intensity: %f\n", PM)
}
