package main

import (
	"awesomeProject/Utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		firstDigit, secondDigit, operator string
		splitInput                        []string
		romansCheck                       = false
	)

	reader := bufio.NewReader(os.Stdin)
	for {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		userInput = strings.Trim(userInput, "\n")
		if len(userInput) < 5 {
			panic("неверный ввод")
		}

		splitInput = strings.Split(userInput, " ")

		if _, ok := Utils.Romans[splitInput[0]]; ok {
			romansCheck = true
		}

		firstDigit, operator, secondDigit = strings.ToUpper(splitInput[0]), splitInput[1], strings.ToUpper(splitInput[2])

		Utils.DigitsCheck(firstDigit, secondDigit, romansCheck, Utils.Romans)

		res := Utils.Result(operator, firstDigit, secondDigit, romansCheck, Utils.Romans)

		fmt.Println(res)

		reader.Reset(reader) //добавил, потому что иногда паниковал без понятной причины
	}
}
