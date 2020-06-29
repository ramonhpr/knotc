package main

import (
	"flag"
	"log"
	"io/ioutil"
	"os"

	"github.com/ramonhpr/knotc/pkg/compiler"
)

var version string

func main() {
	flagVersion := flag.Bool("version", false, "Print compiler version")
	log.SetFlags(0)
	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()
		log.Println(os.Args[0], " filePath outputFolder")
	}
	flag.Parse()
	if *flagVersion {
		log.Println(version)
		os.Exit(0)
	}

	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	filePath := flag.Arg(0)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("file %s not exists", filePath)
	}

	outFile := flag.Arg(1)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Couldn't read file %s. Reason: %s", filePath, err)
	}
	compiler := compiler.NewZephyrCompiler(outFile)
	compiler.Compile(string(data))
}
