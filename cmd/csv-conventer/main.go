package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	originFileName := os.Args[1]
	println("Original file name = ", originFileName)

	file, err := os.Open(originFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	utf8CharSet := charmap.Windows1251.NewDecoder()
	var outFileName = "out_" + originFileName
	outFile, err := os.Create(outFileName)
	if err != nil {
		fmt.Println("Error to create file with name")
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	writer.Write([]byte("Date;Payee;Memo;Outflow\n"))
	var row = 0
	for scanner.Scan() {
		if row == 0 {
			println("Skip first row")
			row++
			continue
		}
		var decodedString = scanner.Text()
		encodedString, err := utf8CharSet.String(decodedString)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(encodedString)
		var splitted = strings.Split(encodedString, ";")
		fmt.Println(splitted[1])
		_, err = writer.Write([]byte(splitted[1] + ";" + ";" + splitted[2] + ";" + splitted[3] + "\n"))
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

}
