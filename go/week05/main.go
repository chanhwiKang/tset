package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	var score string
	if grade >= 90 {
		score = "a grade"
	} else {
		score = "under A grade"
	}
	fmt.Println(input)
	fmt.Println("You will get", score) // undefined: grade

}
