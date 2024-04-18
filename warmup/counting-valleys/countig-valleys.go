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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	if steps == 0 { // if there are no steps, there's no valley
		return 0
	}

	seaLevel := true
	var valley bool
	var valleys int32
	var seq int

	for _, s := range path {
		if s == 'U' { // increase counter if hiker is going up
			seq++
		}
		if s == 'D' { // decrease counter if hiker is going down
			seq--
		}

		if seaLevel && seq == -1 { // if hiker is at sea level, and the counter is -1 it means that he started going down a valley
			valley = true // went into a valley so he's not at sea level anymore
			seaLevel = false
		}

		if valley && seq == 0 { // if he's currently in a valley but the counter got to 0, it means he arrived at sea level
			seaLevel = true
			valley = false
			valleys++ // since hiker just left a valley, let's increase the number of valleys he's gone through
		}
	}

	return valleys
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

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
