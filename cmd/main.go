package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ramonhpr/knotc/pkg/compiler"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		flag.Usage()
		fmt.Println(os.Args[0], " filePath outputFolder")
		os.Exit(1)
	}

	filePath := flag.Arg(0)
	outFile := flag.Arg(1)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	compiler := compiler.NewZephyrCompiler(outFile)
	compiler.Compile(string(data))
}
