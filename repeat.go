package iteration

import "strings"

func Repeat(input string, repeatCount int) (output string) {
	//	for i := 0; i < repeatCount; i++ {
	//		output += input
	//	}
	//	return output
	return strings.Repeat(input, repeatCount)
}
