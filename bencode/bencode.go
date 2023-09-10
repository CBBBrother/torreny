package bencode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)


func printWithTab(tabSize int, content string) {
    log.Print(strings.Repeat("\t", tabSize), " " , content)
}

func readInt(content []byte, delim byte) (value int, size int) {
    for i, v := range content {
        if v == delim {
            value, _ = strconv.Atoi(string(content[:i]))
            size = i
            return
        }
    }

    return 0, -1
}


func ParseBencode(content []byte) {
    tabSize := 0

    for i := 0; i < len(content); i++ {
        switch content[i] {
        case 'd':
            printWithTab(tabSize, "d")
            tabSize++
        case 'i':
            value, size := readInt(content[i + 1:], 'e')
            if size == -1 {
                log.Fatal("Error while parse int")
            }
            printWithTab(tabSize, fmt.Sprintf("i = %v %v", size, value))
            i += size + 1
        case 'l':
            printWithTab(tabSize, "l")
            tabSize++
        case 'e':
            tabSize--
            printWithTab(tabSize, "e")
        default:
            value, size := readInt(content[i:], ':')
            if size == -1 {
                log.Fatal("Error while parse bytes size")
            }
            i += size + 1
            printWithTab(tabSize, fmt.Sprintf("bytes = %v - %v", value, string(content[i:i + value])))
            i += value - 1
        }
    }
}