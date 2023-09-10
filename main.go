package main

import (
	"log"
	"os"
    "cbbb/project37/bencode"
)


func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage: ", os.Args[0], " torrent_filename")
    }

    filename := os.Args[1]

    content, ok := os.ReadFile(filename)

    if ok != nil {
        log.Fatal("Can not read file ", filename)
    }

    bencode.ParseBencode(content)
}