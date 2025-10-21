package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blevesearch/bleve"
)

const indexFile = "./index.bleve"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <lookup_term>\n", os.Args[0])
		os.Exit(0)
	}

	index, err := bleve.Open(indexFile)
	if  err != nil {
		panic(err)
	}

	query := bleve.NewQueryStringQuery(os.Args[1])
	result, err := index.Search(bleve.NewSearchRequest(query))
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
