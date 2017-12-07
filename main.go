package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	params := bumpParams{}
	isStdin := false
	head := 1
	flag.IntVar(&params.Part, "part", 0, "which part (zero-indexed) of the version to bump")
	flag.BoolVar(&params.LeftToRight, "ltr", false, "Read parts left-to-right (default is right-to-left)")
	flag.StringVar(&params.Delimiter, "delimiter", ".", "Delimiter (default is .)")
	flag.StringVar(&params.Prefix, "prefix", ".", "Prefix")
	flag.IntVar(&params.Inc, "inc", 1, "How much to bump")
	flag.StringVar(&params.Sort, "sort", "", "Sort asc/desc (applies to stdin only)")
	flag.IntVar(&head, "head", 1, "First n versions from stdin (specify sort order with -sort). Use -1 to return all records")
	//flag.BoolVar(&isStdin, "stdin", false, "Use standard input")
	flag.Parse()
	if params.Delimiter == "" {
		params.Delimiter = "."
	}
	switch params.Sort {
	case "asc":
	case "desc":
	case "":
		params.Sort = "desc"
	default:
		fmt.Println("Invalid flag for sort. Should be asc or desc")
		os.Exit(1)
	}
	isStdin = len(flag.Args()) < 1
	if isStdin {
		rsorted := []Version{}
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
			v, err := toVersion(strings.TrimSpace(line), &params)
			if err != nil {
				if err != errNoPrefix {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				continue
			}
			rsorted = append(rsorted, v)
		}
		if params.Sort == "asc" {
			sort.Sort(Sorted(rsorted))
		} else {
			sort.Sort(RSorted(rsorted))
		}
		for i, version := range rsorted {
			vNew, err := bump(version, params)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if head < 0 || head > i {
				fmt.Println(vNew)
			}
		}
	} else {
		version, err := toVersion(flag.Arg(0), &params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		vNew, err := bump(version, params)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", vNew)
	}
}
