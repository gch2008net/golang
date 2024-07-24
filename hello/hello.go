package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	message, err := greetings.Hello("Gladys")
	if err != nil {
		fmt.Println("err:" + err.Error())
	}
	fmt.Println(message)

	// message, err := greetings.Hello("")
	// if err != nil {
	// 	fmt.Println("err:" + err.Error())
	// }
	// fmt.Println("msg:" + message)

	// // Request a greeting message.
	// message, err := greetings.Hello("李刚")
	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)

	// // A slice of names.
	// names := []string{"张三", "李四", "老六"}

	// // Request greeting messages for the names.
	// messages, err := greetings.Hellos(names)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // If no error was returned, print the returned map of
	// // messages to the console.
	// fmt.Println(messages)
	// test77.Hello222()
}
