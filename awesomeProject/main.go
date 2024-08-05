package main

import (
	"awesomeProject/Utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
		
	reader := bufio.NewReader(os.Stdin)
	
	for {
		romansCheck := false
		
		fmt.Print("Введите значения: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		userInput = strings.TrimSpace(userInput)
		splitInput := strings.Split(userInput, " ")
		if len(splitInput) != 3 {
			panic("неверный ввод")
		}

		if _, ok := Utils.Romans[splitInput[0]]; ok {
			romansCheck = true
		}

		firstDigit, operator, secondDigit := splitInput[0], splitInput[1], splitInput[2]

		Utils.DigitsCheck(firstDigit, secondDigit, romansCheck, Utils.Romans)

		res := Utils.Result(operator, firstDigit, secondDigit, romansCheck, Utils.Romans)

		fmt.Println(res)
	}
}
