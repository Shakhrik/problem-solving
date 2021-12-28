package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var words = map[string][]int{}
	var N int
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the number of lines: ")
	_, err := fmt.Scan(&N)
	checkError(err)
	for i := 1; i <= N; i++ {
		line, err := reader.ReadString('\n')
		checkError(err)

		ws := strings.Split(line, " ")
		mapWrite(ws, words, i)

	}
	displayOutput(words)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func mapWrite(ws []string, words map[string][]int, line int) {
	for i := 0; i < len(ws); i++ {
		words[strings.TrimSpace(ws[i])] = append(words[strings.TrimSpace(ws[i])], line)
	}
}

func displayOutput(words map[string][]int) {
	keys := make([]string, 0)
	for key := range words {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s %v\n", key, words[key])
	}
}
