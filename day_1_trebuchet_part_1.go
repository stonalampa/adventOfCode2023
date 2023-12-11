package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func isNumber(char byte) bool {
// 	return unicode.IsDigit(rune(char))
// }

func main() {
	filePath := "day_1_trebuchet_input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// ***** Solution 1 *****, i think this is wrong and too complex for no reason
	// * trying to optimize up ahead, premature :(
	// sum := 0
	// lineNum := 0

	// for scanner.Scan() {
	// 	lineNum++
	// 	line := scanner.Text()

	// 	firstNum, secondNum := "", ""
	// 	foundFirst, foundSecond := false, false

	// 	for start := 0; start < len(line)-1; start++ {
	// 		end := len(line) - 1 - start
	// 		if (start > end) || (foundFirst && foundSecond) {
	// 			break
	// 		}

	// 		if isNumber(line[start]) && !foundFirst {
	// 			firstNum = string(line[start])
	// 			foundFirst = true
	// 		}

	// 		if isNumber(line[end]) && !foundSecond {
	// 			secondNum = string(line[end])
	// 			foundSecond = true
	// 		}
	// 	}

	// 	if foundFirst && foundSecond {
	// 		num, err := strconv.Atoi(firstNum + secondNum)
	// 		if err != nil {
	// 			fmt.Println("Error converting string to integer:", err)
	// 			return
	// 		}
	// 		sum += num
	// 	}

	// 	if foundFirst && !foundSecond {
	// 		num, err := strconv.Atoi(firstNum + firstNum)
	// 		if err != nil {
	// 			fmt.Println("Error converting string to integer:", err)
	// 			return
	// 		}
	// 		sum += num
	// 	}

	// 	if !foundFirst && foundSecond {
	// 		num, err := strconv.Atoi(secondNum + secondNum)
	// 		if err != nil {
	// 			fmt.Println("Error converting string to integer:", err)
	// 			return
	// 		}
	// 		sum += num
	// 	}
	// }
	// ************

	// ***** Solution 2 *****
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		arrayNumbers := []string{}
		for i := 0; i < len(line)-1; i++ {
			if line[i] >= '0' && line[i] <= '9' {
				fmt.Println(string(line[i]))
				arrayNumbers = append(arrayNumbers, string(line[i]))
			}
		}

		if len(arrayNumbers) >= 1 {
			firstNum := arrayNumbers[0]
			var lastNum string
			if len(arrayNumbers) > 1 {
				lastNum = arrayNumbers[len(arrayNumbers)-1]
			} else {
				lastNum = firstNum
			}

			concatenated := firstNum + lastNum
			num, err := strconv.Atoi(concatenated)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			sum += num
		}
	}
	// ************

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Final Sum:", sum)
}
