package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	rowsLength := len(arr)
	columnsLength := len(arr[0])
	maxVal := int32(-63) // as per constraints, the maximum number can be negative/positive 9, and the hourglass has 7 numbers, so the smallest can be -9*7

	for i := 0; i < rowsLength-2; i++ { // since we will need to access the current position and two further, we need to avoid breaking the matrix limits
		for j := 0; j < columnsLength-2; j++ {
			submatrix := getSubMatrix(arr, i, i+2, j, j+2)
			sum := sumSubArray(submatrix)
			if sum > maxVal { // get the max value between the hourglass sum and the current max value
				maxVal = sum
			}
		}
	}

	return maxVal
}

func getSubMatrix(matrix [][]int32, startRow, endRow, startCol, endCol int) [][]int32 {
	submatrix := make([][]int32, endRow-startRow+1) // create a submatrix with length as the difference between the expected rows
	for i := range submatrix {
		submatrix[i] = make([]int32, endCol-startCol+1)           // create columns with length as the difference between the expected columns
		copy(submatrix[i], matrix[i+startRow][startCol:endCol+1]) // copy to the current row
	}

	return submatrix
}

func sumSubArray(subArr [][]int32) int32 {
	if len(subArr[0]) < 3 || len(subArr) < 3 { // avoid going further because if for some reason the submatrix comes with less rows or columns, we could have a out of bounds exception
		return 0
	}
	a1, a2, a3 := subArr[0][0], subArr[0][1], subArr[0][2] // values from the first row
	b2 := subArr[1][1]                                     // value from the second row
	c1, c2, c3 := subArr[2][0], subArr[2][1], subArr[2][2] // values from the third row

	return a1 + a2 + a3 + b2 + c1 + c2 + c3
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
