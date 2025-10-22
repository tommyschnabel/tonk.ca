package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/blevesearch/bleve"
)

const (
	indexFile    = "./index.bleve"
	staticDir = "../static/"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file_list>\n", os.Args[0])
		os.Exit(1)
	}
	var index bleve.Index

	// Open or create index file
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		mapping := bleve.NewIndexMapping()
		if index, err = bleve.New(indexFile, mapping); err != nil {
			panic(err)
		}

	} else {
		if index, err = bleve.Open(indexFile); err != nil {
			panic(err)
		}
	}

	files := os.Args[1:]
	fmt.Printf("Got %d files to index\n", len(files))
	for _, f := range files {
		fmt.Println(f)
		bytes, err := os.ReadFile(f)
		if err != nil {
			panic(fmt.Sprintf("error reading file %s: %v", f, err))
		}

		targetName := strings.Split(f, "_OCR")[0] + ".pdf"
		targetName = strings.Split(targetName, staticDir)[1]
		_ = index.Delete(targetName)
		index.Index(targetName, string(bytes))

		fmt.Printf("Indexed: %s\n", targetName)
	}
}
