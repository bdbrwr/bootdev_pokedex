package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input_string := scanner.Text()
		cleaned_string := cleanInput(input_string)
		if len(cleaned_string) == 0 {
			continue
		}
		command := cleaned_string[0]
		fmt.Printf("Your command was: %s \n", command)
	}

}
