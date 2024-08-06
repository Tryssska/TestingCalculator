package main

import (
	"awesomeProject/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		IsRomans := false
		fmt.Print("Введите значения: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		userInput = strings.TrimSpace(userInput)
		splitInput := strings.Split(userInput, " ")

		if _, ok := utils.Romans[splitInput[0]]; ok {
			IsRomans = true
		}

		firstDigit, operator, secondDigit := utils.GetDigits(splitInput, IsRomans)

		res := utils.Result(operator, firstDigit, secondDigit, IsRomans)

		fmt.Println(res)
	}
}
