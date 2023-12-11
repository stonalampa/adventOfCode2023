package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	filePath := "day_2_cube_conundrum_input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		validGame := true
		line := scanner.Text()

		startIndex := strings.Index(line, "Game ")
		if startIndex == -1 {
			fmt.Println("Game not found")
			return
		}

		endIndex := strings.Index(line[startIndex:], ":")
		if endIndex == -1 {
			fmt.Println(": not found")
			return
		}

		gameId, err := strconv.Atoi(line[startIndex+len("Game ") : startIndex+endIndex])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		parts := strings.Split(line[endIndex+2:], ";")
		for _, part := range parts {
			wordsAndNumbers := strings.Fields(part)
			partMap := make(map[string]int)
			for i := 1; i < len(wordsAndNumbers); i += 2 {
				word := strings.TrimRight(wordsAndNumbers[i], ",")
				number, err := strconv.Atoi(wordsAndNumbers[i-1])
				if err != nil {
					fmt.Println("Error converting string to integer:", err)
					return
				}
				partMap[word] = number
			}
			if partMap["red"] > redCubes || partMap["green"] > greenCubes || partMap["blue"] > blueCubes {
				fmt.Println("Game", gameId, "is invalid")
				validGame = false
			}
		}

		if validGame {
			fmt.Println("Game", gameId, "is valid")
			sum += gameId
		}
	}

	fmt.Println("Sum of valid games:", sum)
}
