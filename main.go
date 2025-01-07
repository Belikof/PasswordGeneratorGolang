package main

import (
	"fmt"
	"math/rand"
	"time"
)

var addLetters string
var addNumbers string
var addCharacters string
var lenghtString int

func randomString(lenght int) string {
	charsetLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumbers := "0123456789"
	charsetCharacters := "!@£$%^&*(+"

	result := make([]byte, lenght)
	for i := range result {
		if addLetters == "yes" && addNumbers == "yes" && addCharacters == "yes"{
			result[i] = (charsetLetters + charsetNumbers + charsetCharacters)[rand.Intn(len(charsetLetters + charsetNumbers + charsetCharacters))]
		} else if addLetters == "yes" && addNumbers == "yes" && addCharacters == "no" {
			result[i] = (charsetLetters + charsetNumbers)[rand.Intn(len(charsetLetters + charsetNumbers))]
		} else if addLetters == "yes" && addNumbers == "no" && addCharacters == "no" {
			result[i] = charsetLetters[rand.Intn(len(charsetLetters))]
		} else if addLetters == "no" && addNumbers == "yes" && addCharacters == "no" {
			result[i] = charsetNumbers[rand.Intn(len(charsetNumbers))]
		} else if addLetters == "no" && addNumbers == "yes" && addCharacters == "yes" {
			result[i] = (charsetNumbers + charsetCharacters)[rand.Intn(len(charsetNumbers + charsetCharacters))]
		} else if addLetters == "yes" && addNumbers == "no" && addCharacters == "yes" {
			result[i] = (charsetLetters + charsetCharacters)[rand.Intn(len(charsetLetters + charsetCharacters))]
		} else if addLetters == "no" && addNumbers == "no" && addCharacters == "yes" {
			result[i] = charsetCharacters[rand.Intn(len(charsetCharacters))]
		} else {
			fmt.Println("No way. You need to choose something.")
		}
		
	}
	return string(result)
}


func main() {
	fmt.Println("Welcome to the Password Generator!")
	fmt.Println("Enter the desired password length:")
	fmt.Scan(&lenghtString)
	fmt.Println("Include letters? (yes/no):")
	fmt.Scan(&addLetters)
	fmt.Println("Include numbers? (yes/no):")
	fmt.Scan(&addNumbers)
	fmt.Println("Include special characters? (yes/no):")
	fmt.Scan(&addCharacters)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomString(lenghtString))
}