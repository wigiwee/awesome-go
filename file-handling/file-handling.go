package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// reference : https://www.honeybadger.io/blog/comprehensive-guide-to-file-operations-in-go/

type Config struct {
	DBHost        string `json:"database_host"`
	DBPort        int    `json:"database_port"`
	DBUsername    string `json:"database_username"`
	DBPassword    string `json:"database_password"`
	ServerPort    int    `json:"server_port"`
	ServerDebug   bool   `json:"server_debug"`
	ServerTimeout int    `json:"server_timeout"`
}

func main() {
	filepath := "readme.txt"

	var fileContent string
	fileContent = readEntireFile(filepath)
	fmt.Println(fileContent)

	fmt.Println(readLineByLine("readme.txt"))

	fmt.Println(readJson("db-config.json"))

	fmt.Println(readcsv("carprices.csv"))

	println("reading bytes")

	fmt.Println(readBytes("readme.txt", 15))

	newfile := createFile("abc.txt")
	newfile.Close()

	writeToFile("abc.txt")

	appendToFile("abc.txt", "\nappnding this to the file")
}

func readEntireFile(filepath string) string {
	if len(filepath) > 0 {
		data, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Println("File content \n", string(data))
			return ""
		}
		return string(data)
	} else {
		fmt.Println("filepath cannot be a nil string")
		return ""
	}
}

func readLineByLine(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error while opening file : ", err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // scanner.Text() reads the file line by line, Text() suggests the file to be read in text
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error occured while reading file ", err)
	}

	file.Close()
	return lines
}

func readJson(filepath string) Config {
	var config Config
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error occured while encoding json file ", err)
	}

	file.Close()
	return config
}

func readcsv(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file ", err)
	}
	csvReader := csv.NewReader(file)
	// firstRecord, err := csvReader.Read()	//reads records line by line & return 1D slice
	allRecords, err := csvReader.ReadAll() //reads all the records at once & return 2D slice

	file.Close()
	return allRecords
}

func readBytes(filepath string, bytesToRead int64) []byte {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file ", err)
	}

	data := make([]byte, bytesToRead)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("error reading the file ", err)
	}

	return data
}

func createFile(filepath string) *os.File {
	file, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")

	return file

}

/*
   os.O_RDONLY: Opens the file as read-only. The file must exist.

   os.O_WRONLY: Opens the file as write-only. If the file exists, its contents are truncated. If it doesn't exist, a new file is created.

   os.O_RDWR: Opens the file for reading and writing. If the file exists, its contents are truncated. If it doesn't exist, a new file is created.

   os.O_APPEND: Appends data to the file when writing. Writes occur at the end of the file.

   os.O_CREATE: Creates a new file if it doesn't exist.

   os.O_EXCL: Used with O_CREATE, it ensures that the file is created exclusively, preventing creation if it already exists.

   os.O_SYNC: Open the file for synchronous I/O operations. Write operations are completed before the call returns.

   os.O_TRUNC: If the file exists and is successfully opened, its contents are truncated to zero length.

   os.O_NONBLOCK: Opens the file in non-blocking mode. Operations like read or write may return immediately with an error if no data is available or the operation would block.

*/

func writeToFile(filepath string) bool {

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error opening file ", err)
		return false
	}

	data := "Writing this to abc.txt"
	_, err = file.WriteString(data)

	if err != nil {
		fmt.Println("error writing to file")
		return false
	}
	return true
}

func appendToFile(filepath string, content string) bool {

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error opening file ", err)
		return false
	}

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("error writing/appending to the file ", err)
		return false
	}
	return true

}
