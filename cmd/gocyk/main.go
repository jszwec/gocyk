package main

import (
	"flag"
	"log"
	"os"

	"github.com/jszwec/gocyk/cyk"
	"github.com/jszwec/gocyk/parser"
)

var (
	inputFile  = flag.String("input", "", "file path to Chomsky reduced form grammar")
	outputFile = flag.String("output", "table.html", "output file containing CYK table")
	word       = flag.String("word", "", "word you are willing to check")
)

func main() {
	flag.Parse()
	if *inputFile == "" || *word == "" {
		flag.PrintDefaults()
		return
	}
	g, err := parser.Parse(*inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	ct := cyk.Cyk(g, *word)
	if err = ct.ToFile(*outputFile); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return
}
