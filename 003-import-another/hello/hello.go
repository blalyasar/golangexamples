package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)
func main(){
	log.SetPrefix("greetings")
	// log.setFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	// message, err := greetings.Hello("Glady")
	if err != nil {
		log.Fatal(err)
	}
	// get greeting message and print it 
	// message := greetings.Hello("Gladys")

	fmt.Println(messages)
}
