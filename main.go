package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// s := len(sc.Text())

	fmt.Fprintln(os.Stderr, "エラー")
	fmt.Fprintln(os.Stdout, "標準")
	os.Exit(3)
}
