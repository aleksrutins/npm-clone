package prompt

import "fmt"

// YesNo : Prompt user for y/n
func YesNo(message string, deflt byte) bool {
	var promptStr string
	if deflt == 'y' {
		promptStr = message + " [Yn] "
	} else {
		promptStr = message + " [yN] "
	}
	fmt.Print(promptStr)
	var result string
	fmt.Scan(result)
	if result == "y" || result == "Y" {
		return true
	} else if result == "n" || result == "N" {
		return false
	} else {
		if deflt == 'y' {
			return true
		} else {
			return false
		}
	}
}
