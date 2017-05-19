package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0x75960/fid"
	unqfy "github.com/0x75960/unqfy/lib"
)

var targetDir string

func init() {
	if len(os.Args) != 2 {
		fmt.Printf("\nusage:\n\t %s <target dir path>\n\n", os.Args[0])
		os.Exit(-1)
	}

	targetDir = os.Args[1]
}

func main() {

	all, err := unqfy.Listup(targetDir)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(fid.MarkdownTableTitle())

	for _, file := range all {

		id, err := fid.NewFileIdentifier(file)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(id.MarkdownTableRow())

	}

}
