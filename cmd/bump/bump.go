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
	isStdin := false
	head := 1
	flag.IntVar(&params.Part, "part", 0, "which part (zero-indexed) of the version to bump")
	flag.BoolVar(&params.LeftToRight, "ltr", false, "Read parts left-to-right (default is right-to-left)")
	flag.StringVar(&params.Delimiter, "delimiter", ".", "Delimiter (default is .)")
	flag.StringVar(&params.Prefix, "prefix", ".", "Prefix")
	flag.IntVar(&params.Amount, "inc", 1, "How much to bump")
	flag.StringVar(&params.Sort, "sort", "", "Sort asc/desc (applies to stdin only)")
	flag.IntVar(&head, "head", 1, "First n versions from stdin (specify sort order with -sort). Use -1 to return all records")
	flag.BoolVar(&isStdin, "stdin", false, "Use standard input")
	flag.Parse()
	if params.Delimiter == "" {
		params.Delimiter = "."
	}
	if isStdin {
		rsorted := []bump.Version{}
		stdin := bufio.NewReader(os.Stdin)
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
		}
		if params.Sort == "asc" {
			sort.Sort(bump.Sorted(rsorted))
		} else {
			sort.Sort(bump.RSorted(rsorted))
		}
		for i, version := range rsorted {
			vNew, err := bump.Bump(version, params)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if head < 0 || head > i {
				fmt.Println(vNew)
			}
		}
	} else {
		v := strings.TrimSpace(flag.Arg(0))
		if v == "" {
			//return "", bump.ErrNoVersionSupplied
			fmt.Println("No version supplied")
			os.Exit(1)
		}
		version, err := bump.ToVersion(v, params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		vNew, err := bump.Bump(version, params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", vNew)
	}
}
