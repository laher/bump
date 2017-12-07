package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/laher/bump"
)

func main() {
	params := bump.BumpParams{}
	head := 1
	flag.IntVar(&params.Part, "part", 0, "which part (zero-indexed) of the version to bump")
	flag.BoolVar(&params.LeftToRight, "ltr", false, "Read parts left-to-right (default is right-to-left)")
	flag.StringVar(&params.Delimiter, "delimiter", ".", "Delimiter (default is .)")
	flag.IntVar(&head, "head", 1, "Highest version only. Use -1 to return all records")
	flag.Parse()
	stdin := bufio.NewReader(os.Stdin)
	rsorted := bump.RSorted{}
	for {
		line, err := stdin.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			break
		}
		v, err := bump.ToVersion(strings.TrimSpace(line), params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		rsorted = append(rsorted, v)
		//fmt.Printf("v: %+v\n", v)
	}
	sort.Sort(rsorted)

	for i, version := range rsorted {
		if head < 0 || head > i {
			fmt.Println(version.ToString(params))
		}
	}
}
