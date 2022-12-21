package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pathinfo() string {
	fmt.Println("What is the path of the file")
	inputreader := bufio.NewReader(os.Stdin)
	path, _ := inputreader.ReadString('\n')
	path = strings.TrimSuffix(path, "\n")
	return path
}

func spars(file string) {
	bs, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	failure := strings.Split(string(bs), "\n")
	for _, event := range failure {
		wack := strings.Contains(event, "error")
		if wack {
			fmt.Println(event)
		}
	}
	// fmt.Println(failure)
}
