package main

import "fmt"

//Struct for contact info
type contactInfo struct {
	email   string
	zipCode int
}

//struct for person
type person struct {
	firstName string
	lastName  string
	//for embedded struct we can use: only type as well
	//contact   contactInfo
	contactInfo //the variable name would be same i.e. contactInfo
}

func main() {
	//Creating new object --> could be done ->
	// name := person{"Mayank", "Gupta"}
	//name := person{firstName: "Mayank", lastName: "Gupta"}
	//For embedding struct like above example

	//Last line should have a comma for multi line struct
	personInfo := person{
		firstName: "Mayank",
		lastName:  "Gupta",
		contactInfo: contactInfo{
			email:   "mayank****@gmail.com",
			zipCode: 110062,
		},
	}

	//'&variable' give the memory address of the value this variable is pointing to
	// personPointer --> is memory address now
	/*personPointer := &personInfo
	personPointer.updateName("Mayank1989") */
	personInfo.updateName("Mayank198912") //the above is handled by go [shortcut with go]
	personInfo.print()

	//Trying the same with slice
	slice := []string{"Hi", "How", "are", "you"}
	/* This happens as the slice internal structure -> pointer to head of array, capacity &, length
	* so when we update the slice it actually updates the underlying array even go makes copy of slice
	* similar happens with: slice, maps, channels, pointers &, functions [called Reference Types]
	 */
	updateSlice(slice) //for slices -> Go does not require pointers.
	fmt.Println(slice)
}

//Reciever function for person struct
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

/*Updating the person [with p person it is pass by value so the value is not
updated in the passed struct]
to update we need to pass as: p *person
Go does a pass by value but unlike java it does not pass the reference of the object as value
so in Go we pass:  &person --> which passes the address of the struct in memory
***Whenever we person go creates a copy of the struct &, copy is made available to the update function
hence the change is never propogated to jim. */
func (p *person) updateName(newName string) {
	// '*pointer' -> give me the value this memory address is pointing at
	(*p).firstName = newName
}

func updateSlice(slice []string) {
	slice[0] = "Bye"
}
