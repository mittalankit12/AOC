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
		fmt.Println("error")
		return
	}

	scanner := bufio.NewScanner(file)

	var games []string

	for scanner.Scan() {
		game := scanner.Text()
		games = append(games, game)
	}

	var sum int = 0

	for _, game := range games {
		sum = sum + powerGame(game)
	}
	fmt.Println(sum)

}

func powerGame(s string) int {
	splitColon := strings.Split(s, ":")
	sets := strings.Split(strings.TrimSpace(splitColon[1]), ";")
	red := 0
	green := 0
	blue := 0
	for _, set := range sets {
		balls := strings.Split(strings.TrimSpace(set), ",")
		for _, ball := range balls {
			item := strings.Split(strings.TrimSpace(ball), " ")
			if item[1] == "red" {
				num, _ := strconv.Atoi(item[0])
				if num > red {
					red = num
				}
			}
			if item[1] == "green" {
				num, _ := strconv.Atoi(item[0])
				if num > green {
					green = num
				}
			}
			if item[1] == "blue" {
				num, _ := strconv.Atoi(item[0])
				if num > blue {
					blue = num
				}
			}
		}
	}
	return red * green * blue
}
