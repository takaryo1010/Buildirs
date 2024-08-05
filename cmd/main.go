package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/takaryo1010/buildirs/pkg/buildirs"
)

var (
	path        string
	NewFileName string
	Num         int
	Nums        []int
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: buildirs NewFileName Num \n")
		fmt.Fprintln(os.Stderr, "The number of "+strconv.Quote("*")+" must match the number of numbers")
		// fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
// コメント
	flag.Parse()

	path = "."
	NewFileName = flag.Arg(0)
	for i := 1; i < len(flag.Args()); i++ {
		Num, err := strconv.Atoi(flag.Arg(i))
		Nums = append(Nums, Num)
		if err != nil {
			fmt.Printf("Num = int\n")
			os.Exit(1)
		}
	}
	if len(Nums) != strings.Count(NewFileName, "*") {
		text := "The number of " + strconv.Quote("*") + " does not match the number of digits"
		fmt.Println(text)
		os.Exit(1)
	}

	NewFileNames := strings.Split(NewFileName, "/")

	err := buildirs.Buildirs(path, NewFileNames, Nums, -1)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
