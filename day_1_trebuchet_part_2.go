package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	filePath := "day_1_trebuchet_input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		arrayNumbers := []string{}
		localStringNum := ""
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				arrayNumbers = append(arrayNumbers, string(line[i]))
				localStringNum = ""
			} else {
				localStringNum += string(line[i])
				// check all substring if any match
				for j := 0; j < len(localStringNum)-1; j++ {
					for z := j + 1; z <= len(localStringNum); z++ {
						substring := localStringNum[j:z]
						if val, exists := numbersMap[substring]; exists {
							fmt.Println("EVO GA", localStringNum)
							arrayNumbers = append(arrayNumbers, val)

						}
					}
				}
			}
		}

		fmt.Println("ARRAY NUMBERS", arrayNumbers)
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

			fmt.Println("Num:", num)
			sum += num
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Final Sum:", sum)
}
