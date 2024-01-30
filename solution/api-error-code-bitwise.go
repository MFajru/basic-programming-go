package solution

import "fmt"

func errorBitwise(input int) []string {
	errorCodeMap := map[int]string{
		8: "You are not authorized to proceed with the input",
		4: "The server is overloaded by too much traffic",
		2: "The server encounters internal error",
		1: "Incorrect input",
	}

	var errorKeys []int

	for key := range errorCodeMap {
		errorKeys = append(errorKeys, key)
	}

	tempInput := input
	var messages []string

	if tempInput == 0 {
		messages = append(messages, "No Error")
		return messages
	}

	// if tempInput

	for _, errorKey := range errorKeys {
		if errorKey <= tempInput {
			tempInput = tempInput ^ errorKey
			messages = append(messages, errorCodeMap[errorKey])
		}
	}

	return messages
}

// 16 = 10000 = 10000 1000
// 8 =  01000 = 10111 0010
//              10000
// 0011
// 0101
// 0110

func ErrorCode() {
	var input int

	fmt.Print("Input error code: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Printf("Invalid input, %s. Please input valid combination (i.e. 10 (8+2))", err)
	}
	errorMessages := errorBitwise(input)
	fmt.Println(errorMessages)
	fmt.Println(2 & 8)
}
