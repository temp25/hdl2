package utils

import (
    "fmt"
    "strings"
)

func ParseM3u8Content(m3u8Content string) {
    
    var infoArray map[string]int
    for index, line := range strings.Split(m3u8Content, "\n") {
        //fmt.Printf("\nindex %d, line %s\n",index, line)
        if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
            lineM strings.Replace(str, "#EXT-X-STREAM-INF:", "", -1)
        } else {
            
        }
    }
    
}