package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func fileProcessor(data []byte, lineChan chan string) {
        scanner := bufio.NewScanner(bytes.NewReader(data))
        for scanner.Scan() {
                lineChan <- scanner.Text() // send every line to channel
        }
        close(lineChan)
}

func resultWriter(resultChan chan string) {
        for line := range resultChan {
                // perform processing
                processedLine := strings.ToUpper(line)
                resultChan <- processedLine // send processed line to channel
        }

        close(resultChan)
}

func main() {
	// read file contents
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create channels for data flow
	lineChan := make(chan string) // channel for lines to process
	resultChan := make(chan string) // channel for processed results

	// spawn go routines
	go fileProcessor(data, lineChan) // feeds lines to channel
	go resultWriter(resultChan) // writes procesed results

	// loop to receive processed data
	for processedLine := range resultChan {
		fmt.Println("Processed line:", processedLine)
	}
}

