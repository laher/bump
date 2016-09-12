package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/laher/bump"
)

func main() {
	params := bump.BumpParams{}
	isStdin := false
	flag.IntVar(&params.Part, "part", 0, "which part (zero-indexed) of the version to bump")
	flag.BoolVar(&params.LeftToRight, "ltr", false, "Read parts left-to-right (default is right-to-left)")
	flag.StringVar(&params.Delimiter, "delimiter", ".", "Delimiter (default is .)")
	flag.BoolVar(&isStdin, "stdin", false, "Use standard input")
	flag.Parse()
	v := flag.Arg(0)
	if isStdin {
		stdin := bufio.NewReader(os.Stdin)
		line, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		params.V = line
	} else {
		params.V = v
	}
	vNew, err := bump.Bump(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s\n", vNew)
}
