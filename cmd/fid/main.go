package main

import (
	"fmt"
	"log"
	"os"

	fid "github.com/0x75960/fid"
)

var targetPath string

func init() {
	if len(os.Args) != 2 {
		fmt.Printf("\nusage:\n\t %s <target path>\n\n", os.Args[0])
		os.Exit(-1)
	}

	targetPath = os.Args[1]
}

func main() {
	fid, err := fid.NewFileIdentifer(targetPath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(fid.MarkdownTable())
}
