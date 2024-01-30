//With Slices

package solution

func ApiErrorCodeV2(errorInput int) []string {
	errorCode := []int{8, 4, 2, 1}
	errorMess := []string{
		"You are not authorized to proceed with the input",
		"The server is overloaded by too much traffic",
		"The server encounters internal error",
		"Incorrect input",
	}
	sliceMess := []string{}
	tempErrorVal := errorInput

	if errorInput == 0 {
		sliceMess = append(sliceMess, "No Error")
	}

	for i := 0; i < len(errorCode); i++ {
		tempErrorVal -= errorCode[i]
		if tempErrorVal < 0 {
			tempErrorVal += errorCode[i]
		} else if tempErrorVal > 0 {
			sliceMess = append(sliceMess, errorMess[i])
		} else {
			sliceMess = append(sliceMess, errorMess[i])
		}
	}

	return sliceMess
}
