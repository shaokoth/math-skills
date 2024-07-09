# Statistical Analysis Program

## Overview

This program reads a list of float64s from a file and calculates the mean, median, variance, and standard deviation. It then prints these statistical values to the console.

## Features

* Reads float64 values from a file of type string
* calculates the mean
* Calculates the Median
* Calculates the Variance
* Calculates the Standard deviation

## Prerequisites

1. To run this program make sure you have GO 1.16 or above installed on your system. You can download and install Go from [golang.org](https://golang.org/dl/).

2. Clone this repository or download source code.
https://learn.zone01kisumu.ke/git/shaokoth/math-skills.git

## Usage

Ensure you have a file (e.g., data.txt) in the same directory as the program. This file should contain the population data whose statistics we are interested in. 

Run the program with the default data.txt file:
    go run . main.go data.txt
    
## Code Structure

main.go: The main file containing the code logic.
read_math_skills.go: function where the file is open and read.
mean.go: calculates the average of the population data.
median.go: calculates the median value from the population data.
variance.go: calculates the varaince of the population data.
standarddeviation.go: calculates the standard deviation. This is found from the squareroot of the variance.

## Error Handling

The program includes basic error handling for:

    * Invalid file names or paths.
    * Errors during file reading.
    * Unsupported characters not defined in the data file.
    * Empty files
    
## Other Innformation
This project is maintained by 
@ShadrackOkoth
