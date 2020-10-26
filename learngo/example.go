// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

// func main() {
// 	// totalLenght, up := lenAndUpper("woung")
// 	// fmt.Println(totalLenght, up)

// 	// result := superAdd(1, 2, 3, 4, 5, 6, 7)
// 	// fmt.Println(result)

// 	// fmt.Println(canIDrink(16))

// 	// fmt.Println(canIDrinkSwitch(18))

// 	// Low-Level Programming : Memory Address, Pointers

// 	// fmt.Println(&a, &b) // & => variable's memory address in Go

// 	// a := 2
// 	// b := &a
// 	// a = 5
// 	// fmt.Println(*b) // * => pointers in Go

// 	// a := 2
// 	// b := &a
// 	// *b = 20
// 	// fmt.Println(a)

// 	// a := 3
// 	// b := &a
// 	// *b = 2020
// 	// fmt.Println(&b, &a) // not equal
// 	// fmt.Println(b, &a)  // equal

// 	// Array
// 	//names := [5]string{"woung", "woung1", "woung2"}
// 	// names := []string{"woung", "woung1", "woung2"} // no limit on Array
// 	// names = append(names, "flynn")                 //functon append() returns modified array

// 	// // names[5] = "alalal"
// 	// fmt.Println(names)

// 	// Map
// 	// woung := map[string]string{"name": "woung", "age": "18"} // Map in Go => {key:value} but only one data type to define

// 	// for key, value := range woung {
// 	// 	fmt.Println(key, value)
// 	// }

// 	//Struct
// 	favFood := []string{"kimchi", "cheese"}
// 	woung := person{"woungsub", 18, favFood}
// 	woung2 := person{name: "woungsub", age: 18, favFood: favFood}
// 	fmt.Println(woung)
// 	fmt.Println(woung.name)
// 	fmt.Println(woung2)

// }

// func lenAndUpper(name string) (lenght int, uppercase string) { //naked func =>
// 	defer fmt.Println("I'm done") // similar with javascript callback function!!
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func superAdd(numbers ...int) int {
// 	for index, number := range numbers {
// 		fmt.Println(index, number)
// 	}
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func canIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 { // variable expression => variable for only if script
// 		return false
// 	}
// 	return true
// 	// fmt.Println(koreanAge) //false
// }

// func canIDrinkSwitch(age int) bool {
// 	switch koreanAge := age + 2; koreanAge {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }