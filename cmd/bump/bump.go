package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/laher/bump"
)

func main() {
	params := bump.BumpParams{}
	flag.IntVar(&params.Part, "dot", 0, "how many dots")
	flag.BoolVar(&params.LeftToRight, "ltr", false, "Left to right")
	flag.Parse()
	v := flag.Arg(0)
	params.V = v
	vNew, err := bump.Bump(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s\n", vNew)
}
