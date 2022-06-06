package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := len(sc.Text())

	fortune := [...]string{"大吉", "中吉", "小吉", "吉", "末吉", "凶", "大凶"}

	dt := time.Now()
	u := int(dt.Unix() % 20)
	fmt.Println(fortune[((s+u)*3-4)%7])
}
