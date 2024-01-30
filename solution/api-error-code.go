package solution

import (
	"sort"
)

func sortKeyMap(errorCodeMap map[int]string) []int {
	errorCode := []int{}

	for k := range errorCodeMap {
		errorCode = append(errorCode, k)
	}
	sort.Ints(errorCode)

	return errorCode
}

func ApiErrorCode(input int) []string {

	errorCodeMap := map[int]string{
		1: "Incorrect input",
		2: "The server encounters internal error",
		4: "The server is overloaded by too much traffic",
		8: "You are not authorized to proceed with the input",
	}

	var sliceMessages []string
	sortedErrCode := sortKeyMap(errorCodeMap)
	tempValue := input

	if input == 0 {
		sliceMessages = append(sliceMessages, "No Error")
	}

	for i := len(sortedErrCode) - 1; i >= 0; i-- {
		tempValue -= sortedErrCode[i]
		if tempValue < 0 {
			tempValue += sortedErrCode[i]
		} else if tempValue > 0 {
			sliceMessages = append(sliceMessages, errorCodeMap[sortedErrCode[i]])
		} else if tempValue == 0 {
			sliceMessages = append(sliceMessages, errorCodeMap[sortedErrCode[i]])
			break
		}
	}

	return sliceMessages
}
