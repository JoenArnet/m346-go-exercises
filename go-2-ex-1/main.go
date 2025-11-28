package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName  string
	lastname   string
	dayOfBirth byte
}

// TODO: declare a structure for birth date
type brithDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	brithDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			firstName:  "Joen",
			lastname:   "Arnet",
			dayOfBirth: 10,
		},
		brithDate: brithDate{
			DayOfBirth:   10,
			MonthOfBirth: 9,
			YearOfBirth:  2008,
		},
		NumberOfSiblings: 1,   // TODO: adjust
		ZodiacSign:       'J', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
