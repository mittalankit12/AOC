package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	comstr := ""
	for _, char := range s {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		comstr += string(char)
		for i, num := range nums {
			if strings.Contains(comstr, num) {
				return i + 1
			}
		}
	}
	return 0
}

func lastInt(s string) int {
	nums := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
	comstr := ""
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		comstr += string(char)
		for i, num := range nums {
			if strings.Contains(comstr, num) {
				return i + 1
			}
		}
	}
	return 0
}
