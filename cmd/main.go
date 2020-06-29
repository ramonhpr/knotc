package main

import (
	"flag"
	"log"
	"io/ioutil"
	"os"

	"github.com/ramonhpr/knotc/pkg/compiler"
)

var version string

func usage() {
	log.Println(os.Args[0], "[options] filePath outputFolder")
	log.Println("Options:")
	flag.PrintDefaults()
}

var flagVersion bool
func init() {
	flag.BoolVar(&flagVersion, "version", false, "Print compiler version")
	flag.BoolVar(&flagVersion, "v", false, "Print compiler version (shorthand)")
	flag.Usage = usage
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	if flagVersion {
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
