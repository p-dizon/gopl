package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"./shafactory"
)

var shaSumFlag = flag.Int("sum", 256, "256, 384, or 512")

func main() {
	flag.Parse()
	shaSumFunc, err := shafactory.GetShaSumFunc(*shaSumFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("> ")
	userInput := getUserInput()
	fmt.Printf("%x\n", shaSumFunc(userInput))
}

func getUserInput() []byte {
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInput = strings.TrimSpace(userInput)
	return []byte(userInput)
}
