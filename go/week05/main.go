package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if input >= 90 { // missmatched types string and untyped int
		grade := "a Grade"
	} else {
		grade := "under A grade"
	}
	fmt.Println(input)

}
