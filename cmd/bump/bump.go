package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/laher/bump"
)

func main() {
	partToBump := flag.Int("dot", 0, "how many dots")
	leftToRight := flag.Bool("ltr", false, "Left to right")
	flag.Parse()
	v := flag.Arg(0)
	vNew, err := bump.Bump(v, *partToBump, *leftToRight)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s\n", vNew)
}
