package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("errors")
		return
	}

	scanner := bufio.NewScanner(file)

	var games []string

	for scanner.Scan() {
		game := scanner.Text()
		games = append(games, game)
	}

	var sum int = 0

	for i, game := range games {
		isPossible := possibleGame(game)
		if isPossible {
			sum = sum + i + 1
		}
	}

	fmt.Println(sum)

}

type Cube struct {
	red   int
	green int
	blue  int
}

func possibleGame(s string) bool {
	cubes := Cube{12, 13, 14}
	splitColon := strings.Split(s, ":")
	sets := strings.Split(strings.TrimSpace(splitColon[1]), ";")

	for _, set := range sets {
		balls := strings.Split(strings.TrimSpace(set), ",")
		for _, ball := range balls {
			item := strings.Split(strings.TrimSpace(ball), " ")
			if item[1] == "red" {
				num, _ := strconv.Atoi(item[0])
				if num > cubes.red {
					return false
				}
			}
			if item[1] == "green" {
				num, _ := strconv.Atoi(item[0])
				if num > cubes.green {
					return false
				}
			}
			if item[1] == "blue" {
				num, _ := strconv.Atoi(item[0])
				if num > cubes.blue {
					return false
				}
			}
		}
	}
	return true
}
