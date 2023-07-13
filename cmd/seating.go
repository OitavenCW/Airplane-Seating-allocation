/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var Groups [][]int
var RemainingFamilySize int
var fileName, familySize string

// seatingCmd represents the seating command
var seatingCmd = &cobra.Command{
	Use:   "seating",
	Short: "This program assigns seats to families based on their size in a venue with a given seating layout and prints the seating information for each family.",
	Long: `
This Go program assigns seats to families based on their size in a venue. 
It reads a seating layout from a file and systematically searches for available seats, 
ensuring that each family is accommodated. 
The program then prints the seating information for each family, indicating the row and seat numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		allocate(fileName, familySize)
	},
}

func init() {
	rootCmd.AddCommand(seatingCmd)

	seatingCmd.Flags().StringVarP(&fileName, "fileName", "n", "", "File to be read")
	if error := seatingCmd.MarkFlagRequired("fileName"); error != nil {
		fmt.Println(error)
	}

	seatingCmd.Flags().StringVarP(&familySize, "familySize", "s", "0", "size of the group to be allocated")
	if error := seatingCmd.MarkFlagRequired("familySize"); error != nil {
		fmt.Println(error)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seatingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seatingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func allocate(FileName, FamilySize string) {

	familySize, err := strconv.Atoi(FamilySize)

	if err != nil {
		log.Fatal(err)
	}

	RemainingFamilySize = familySize

	rows := readFile(FileName)

	for RemainingFamilySize > 0 {
		getSeats(rows, RemainingFamilySize)
	}

	if RemainingFamilySize > 0 {
		fmt.Println("Infelizmente não existem lugares para acomodar toda a sua familia.")
		return
	}

	printResponse()
}

func readFile(fileName string) [][]byte {
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

			Groups = append(Groups, group)

			rows[index] = replaceSeats(line, string(subString), string(bytes.Repeat([]byte("X"), familySize)))
			RemainingFamilySize -= familySize

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
	if len(Groups) != 0 {
		if len(Groups) == 1 {

			if Groups[0][1] == Groups[0][2] {
				fmt.Printf("Você pode se sentar na fileira %d, assento %d.\n", Groups[0][0], Groups[0][1])
			} else {
				fmt.Printf("A sua familia pode se sentar na Fileira %d, do assento %d até o assento %d.\n", Groups[0][0], Groups[0][1], Groups[0][2])
			}

		} else {
			fmt.Println("A sua familia teve que ser dividida, eles podem se sentar nos seguintes locais:")
			for _, seats := range Groups {
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
