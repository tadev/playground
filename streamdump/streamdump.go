package main

import (
	"github.com/mewmew/playground/streamdump/dump"

	"flag"
	"fmt"
	"log"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Println("streamdump [URL]...")
}

func main() {
	flag.Parse()
	for _, rawUrl := range flag.Args() {
		err := dump.Url(rawUrl)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
