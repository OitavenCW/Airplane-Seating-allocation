# Ariplane-Seating-allocation

This Go program reads a seating arrangement file and tries to find suitable seats for a given family size. The program attempts to seat the entire family together in adjacent seats or in the same row if possible. If there are not enough available seats in a row, the family members will be divided into multiple groups to accommodate as many family members together as possible.


## Usage


Ensure you have Go installed on your system.

Download or clone the project.

Compile and run the program using the following command:

```sh
go run main.go <filename> <family_size>
```

Replace <filename> with the path to the seating arrangement file, and <family_size> with a positive whole number of family members you want to seat.

## Input File Format

The input file should contain rows of seats represented by the characters 'D' (denoting an available seat) and 'X' (denoting an occupied seat). The program will search for available seats to accommodate the family members based on their family size.

Example of a seating arrangement file (seating.txt):

```
DXD DXD DDD
XXX DXX DDX
XDD XDD XDX
```

## Output

The program will output the seating information for the family based on the available seats. It will either indicate a single seating group or multiple groups if the family needs to be divided.

Example output:

```
A sua familia pode se sentar na fileira 1, do assento 1 até o assento 4.
```
or

```
A sua familia teve que ser dividida, eles podem se sentar nos seguintes locais:
Fileira 1, do assento 1 até o assento 4.
Fileira 2, do assento 1 até o assento 3.
Fileira 3, assento 1.
```

If no available seating arrangement is found for the entire family, the program will display the following message:

```
Infelizmente não existem lugares para acomodar toda a sua familia.
```

## Implementation Details

The program works by recursively searching for available seating arrangements starting from the largest family size down to 1. It keeps track of the remaining family size that still needs to be seated. Once a suitable seating group is found, it marks those seats as occupied and reduces the remaining family size accordingly.

The program considers different seating layouts and adjusts the seat range accordingly to ensure that the family members can sit together without gaps.

Please note that this program assumes a valid input file format with consistent row lengths and available seat representations.

## Limitations

The program assumes a rectangular seating layout with rows of equal length.

It does not handle irregular seating arrangements or complex seating rules.

