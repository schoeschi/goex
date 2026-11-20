package main

import "fmt"

func main() {
	firstName := "joshua"
	lastName := "dietiker"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	dayOfBirth := 1
	monthOfBirth := 12
	yearOfBirth := 2010
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	numberOfSiblings := 1
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	heightInMeters := 169.0
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	zodiacSign := '\u2650'
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
