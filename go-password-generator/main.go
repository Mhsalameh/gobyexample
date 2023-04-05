package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChar = "!@#$%*_+"
	numberSet = "0123456789"
)

func generatePass(passLength, minSpecialChar, minNum, minUpperCase int){
	var password strings.Builder

	for i := 0; i < minSpecialChar; i++{
		random:= rand.Intn(len(specialChar))
		password.WriteString(string(specialChar[random]))
	}

	for i:=0 ; i < minNum; i++{
		random:= rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	for i:=0; i<minUpperCase ; i++{
		random := rand.Intn(len(upperCase))
		password.WriteString(string(upperCase[random]))
	}

	remainingLength := passLength - minSpecialChar - minNum - minUpperCase
	for i:=0; i<remainingLength; i++{
		random:=rand.Intn(len(lowerCase))
		password.WriteString(string(lowerCase[random]))
	}
	inRune:= []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i],inRune[j] = inRune[j],inRune[i]
	})
	
	fmt.Println(string(inRune))

}

func main(){
	var passLength int
	var minSpecialChar int
	var minNum int
	var minUpperCase int
	fmt.Print("password length: ")
	fmt.Scanln(&passLength)
	fmt.Print("Number of special characters: ")
	fmt.Scanln(&minSpecialChar) 
	fmt.Print("Number of number characters: ")
	fmt.Scanln(&minNum) 
	fmt.Print("Number of upper case characters: ")
	fmt.Scanln(&minUpperCase) 
 
	generatePass(passLength,minSpecialChar,minNum,minUpperCase)
}