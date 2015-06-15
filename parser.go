package main

import (
	"fmt"
	"os"

	"github.com/dotabuff/yasha"
  "github.com/davecgh/go-spew/spew"
)

const MAX_COORDINATE float64 = 16384

// Example of printing any updates to hero coordinates

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a .dem file as argument")
	}

	for _, path := range os.Args[1:] {
		parser := yasha.ParserFromFile(path)
		parser.OnEntityPreserved = func(pe *yasha.PacketEntity) {
      spew.Dump(pe)
		}
		parser.Parse()
	}
}

type Coordinate struct {
	X, Y float64
}

func hpFromCell(pe *yasha.PacketEntity) {
  spew.Dump(pe)
}
