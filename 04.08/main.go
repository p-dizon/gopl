package main

import (
	"fmt"
	"os"

	"./charcount"
)

func main() {
	counts := charcount.CharCount(os.Stdin)
	fmt.Println(counts)
}
