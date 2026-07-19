package main

import (
	"blackbox/gemini"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[1;31m>>>\033[0m ")

		prompt, err := reader.ReadString('\n')
		prompt = strings.TrimSpace(prompt)
		if err!=nil{
			fmt.Println("Error reading string.")
			break
		}
		
		
		if prompt == "exit" {
			break
		}
		fmt.Println()
		gemini.CallGemini(prompt)
	}

	
}