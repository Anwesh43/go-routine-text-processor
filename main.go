package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func processText(words []string, ch chan string) {
	wordStr := ""

	for _, w := range words {
		wordStr = fmt.Sprintf("%s %s \n", wordStr, w)
		time.Sleep(1 * time.Second)
		fmt.Println("From go routine 2: Processing 1 line", w)
	}
	ch <- wordStr
}

func writeToFile(lines string, ch chan bool) {
	f, err := os.Create("result.txt")
	if err != nil {
		panic("Error creating file")
		ch <- false
	}
	defer f.Close()
	_, err = f.Write([]byte(lines))
	if err == nil {
		fmt.Println("Successfully written to file ")
		ch <- true
	} else {
		panic("Error writing to file")
		ch <- false
	}

}

func getInputLines(ch chan []string) {
	input := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for true {
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", 1)
		if word == "QUIT" {
			break
		}
		input = append(input, word)
	}
	fmt.Println("Successfully Got Input")
	ch <- input
}

func main() {

	ch1 := make(chan []string)
	ch2 := make(chan string)
	ch3 := make(chan bool)
	go getInputLines(ch1)
	words := <-ch1
	go processText(words, ch2)
	result := <-ch2
	go writeToFile(result, ch3)
	status := <-ch3
	if status {
		fmt.Println("Successfully written to file from a go routine")
	} else {
		fmt.Println("Error writting to a file")
	}
}
