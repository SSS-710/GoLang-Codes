package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LearnCodeOnline.in\n"
	content += "Learning Go file handling.\n"
	content += "Created on: " + time.Now().Format("02-01-2006 15:04:05")

	file, err := os.Create("./mylcogofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
	fileInfo("./mylcogofile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("\nText data inside the file:")
	fmt.Println(string(dataByte))
}

func fileInfo(filename string) {
	info, err := os.Stat(filename)
	checkNilErr(err)

	fmt.Println("\nFile Information:")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Last Modified:", info.ModTime())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
