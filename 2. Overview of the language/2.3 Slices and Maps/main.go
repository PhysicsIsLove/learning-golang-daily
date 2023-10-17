package main

import "log"

func main() {
	names := make(map[string]string)

	names["firstName"] = "Abhishek"
	names["secondName"] = "Anand"

	log.Println(names)


	list := []string{"abc", "def"}
	log.Println("list ", list)
}