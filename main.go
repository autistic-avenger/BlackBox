package main

import (
	"blackbox/gemini"
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Enter The Prompt:")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">>>")
		if !scanner.Scan() {
			break
		}
		
		prompt := scanner.Text()
		if prompt == "exit" {
			break
		}

		gemini.CallGemini(prompt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	
}