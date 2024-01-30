package solution

import (
	"fmt"
	"strings"
)

func RobotTranslator(commands string) {
	var robotDirection string
	var tempCommand string
	var counter int = 1
	wordTime := "time"

	sliceCommands := strings.Split(commands, "")

	for i := 0; i < len(sliceCommands); i++ {

		tempCommand = sliceCommands[i]
		switch sliceCommands[i] {
		case "R":
			robotDirection = "right"
		case "A":
			robotDirection = "advance"
		case "L":
			robotDirection = "left"
		default:
			fmt.Println("Invalid command. Program has stopped")
			return
		}

		if counter > 1 {
			wordTime = "times"
		} else {
			wordTime = "time"
		}

		if i == len(sliceCommands)-1 {
			fmt.Printf("Move %s %d %s\n", robotDirection, counter, wordTime)
			break
		}

		if tempCommand == sliceCommands[i+1] {
			counter++
		} else {
			fmt.Printf("Move %s %d %s\n", robotDirection, counter, wordTime)
			counter = 1
		}

	}
}
