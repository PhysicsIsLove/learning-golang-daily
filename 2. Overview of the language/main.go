// package main

// import (
// 	"fmt"
// 	"log"
// )

// func main() {

// 	var day string = "today"
// 	fmt.Println("Hello world")
// 	fmt.Println("the day is ", day)

// 	nextDay := getNextDay();

// 	fmt.Println("the next day is ", nextDay);

// 	nextDay, nextToNextDay := getNextTwoDays()

// 	fmt.Println("the next two days are", nextDay, "and",  nextToNextDay);

// 	log.Println("the referense of the day is ", &day)
// 	changeUsingPointer(&day)
// 	log.Println("After changing the variable using pointer and reference  ", day)
// }

// func getNextDay() string {
// 	return "tomorrow"
// }

// func getNextTwoDays() (string, string) {
// 	return "tomorrow", "day after tomorrow"
// }

// func changeUsingPointer(s *string){
// 	var newStr string = "yesterday"
// 	*s = newStr;
// }