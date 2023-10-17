It is a good practice to call you main go program as main.go

Every go program has to have al teast one function called main.

fmt stands for format, It is a prt of the standard Go library

the first line in every go program must be a package declaration

Note that we can return something like this, as in more than one return value
`
func getNextTwoDays() (string, string) {
	return "tomorrow", "day after tomorrow"
}
`

To assign th return of any function to another variable we can use the `:=` sign.
Note that because of this we don't need to set the type of the variable, it gets auto assigned to the type returned by the function.


we can get the address (reference to any variable) of any variable using the & operator before it
When you want it to refer to any actual variable, then you put an asterisk infront of it

