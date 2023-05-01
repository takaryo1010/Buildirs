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
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: buildirs PATH NewFileName Num \n")
		// fmt.Fprintln(os.Stderr, "This is my command.")
		// fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	path = flag.Arg(0)
	NewFileName = flag.Arg(1)
	Num, err := strconv.Atoi(flag.Arg(2))
	fmt.Println(path, NewFileName, Num)

	if err != nil {
		fmt.Printf("Num = int\n")
		os.Exit(1)
	}
	NewFileNames := strings.Split(NewFileName, "/")
	err = buildirs.Buildirs(path, NewFileNames, Num)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
