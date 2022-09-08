package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // Syntactic Sugar, doesn't work for constants or if you want to declare the type
	const conferenceTickets = 50      // Go can inference the type
	var remainingTickets uint = 50    // But, we can explicitly specify the type if we want

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) // %T returns the type of the variable

	fmt.Println("Welcome to", conferenceName, "booking application")                                            // Println: Print and add a new line at the end
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets) // Printf: Allow the use of Template String, needs to add line break manually
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// fmt.Println(remainingTickets) // Returns the value
	// fmt.Println(&remainingTickets) // Returns the Pointer: a variable that stores the address of where the variable remainingTickets is stored

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}
