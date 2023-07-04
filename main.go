package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var groups [][]int
var remainingFamilySize int

func main() {

	fileName := os.Args[1]
	familySize, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	remainingFamilySize = familySize

	rows := ReadFile(fileName)

	for remainingFamilySize > 0 {
		getSeats(rows, remainingFamilySize)
	}

	if remainingFamilySize > 0 {
		fmt.Println("Infelizmente não existem lugares para acomodar toda a sua familia.")
		return
	}

	printResponse()

}

func ReadFile(fileName string) [][]byte {
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([][]byte, 0)

	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getSeats(rows [][]byte, familySize int) {

	if familySize <= 0 {
		return
	}

	for index, line := range rows {

		subString := bytes.Repeat([]byte("D"), familySize)
		seatStart := bytes.Index(line, subString) + 1

		if bytes.Contains(line, subString) {
			group := make([]int, 3)

			group[0] = index + 1

			if seatStart > 6 {
				group[1] = seatStart - 2
				group[2] = seatStart + familySize - 3

			} else if seatStart > 3 {
				group[1] = seatStart - 1
				group[2] = seatStart + familySize - 2

			} else {
				group[1] = seatStart
				group[2] = seatStart + familySize - 1
			}

			groups = append(groups, group)

			rows[index] = replaceSeats(line, string(subString), string(bytes.Repeat([]byte("X"), familySize)))
			remainingFamilySize -= familySize

			return

		}
	}

	if familySize >= 1 {
		getSeats(rows, familySize-1)
	}

}

func replaceSeats(line []byte, subString, replacement string) []byte {
	index := bytes.Index(line, []byte(subString))
	if index == -1 {
		return line
	}

	copy(line[index:index+len(replacement)], replacement)
	return line
}

func printResponse() {
	if len(groups) != 0 {
		if len(groups) == 1 {

			if groups[0][1] == groups[0][2] {
				fmt.Printf("Você pode se sentar na fileira %d, assento %d.\n", groups[0][0], groups[0][1])
			} else {
				fmt.Printf("A sua familia pode se sentar na Fileira %d, do assento %d até o assento %d.\n", groups[0][0], groups[0][1], groups[0][2])
			}

		} else {
			fmt.Println("A sua familia teve que ser dividida, eles podem se sentar nos seguintes locais:")
			for _, seats := range groups {
				if seats[1] == seats[2] {
					fmt.Printf("Fileira %d, assento %d.\n", seats[0], seats[1])

				} else {
					fmt.Printf("Fileira %d, do assento %d até o assento %d.\n", seats[0], seats[1], seats[2])
				}
			}
		}
	} else {
		fmt.Println("Por favor entre um numero valido de pessoas.")
	}
}
