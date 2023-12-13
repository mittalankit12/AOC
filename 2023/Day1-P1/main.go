package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error")
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var sum int = 0

	for _, line := range lines {
		fistInt := firstInt(line)
		lastInt := lastInt(line)
		sum = sum + fistInt*10 + lastInt
	}

	fmt.Println(sum)
}

func firstInt(s string) int {
	for _, char := range s {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}

	}
	return 0
}

func lastInt(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
	}
	return 0
}
