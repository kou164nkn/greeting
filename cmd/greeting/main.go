package main

import (
	"os"

	"github.com/kou164nkn/greeting"
)

func main() {
	var g greeting.Greeting

	g.Do(os.Stdout)
}
