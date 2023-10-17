If a variable is starting with a smallletter, then it can only be accessed in the same package
If starting with a capital letter, then can be accesses in other packages also

variables can be declared at the package level also. With that are are accessible everywhere
the scoping is very similar to the scoping of other languages

Golang is not an object oriented programming language

The way to write a function for a struct is this 
    `func (m *User) getFirstName() string {
	return m.FirstName
}`
