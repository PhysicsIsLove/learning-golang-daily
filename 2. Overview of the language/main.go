package main

import "fmt"

func main() {

	var day string = "today"
	fmt.Println("Hello world")
	fmt.Println("the day is ", day)

	nextDay := getNextDay();

	fmt.Println("the next day is ", nextDay);

	nextDay, nextToNextDay := getNextTwoDays()

	fmt.Println("the next two days are", nextDay, "and",  nextToNextDay);
}

func getNextDay() string {
	return "tomorrow"
}

func getNextTwoDays() (string, string) {
	return "tomorrow", "day after tomorrow"
}