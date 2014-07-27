package main

import (
    "fmt"
    "os"
    "datadriver"
    "strings"
)

func main() {
    var path string
    var action string
    var data string

    if len(os.Args) > 2 {
        path = os.Args[2]
        action = os.Args[1]
    }
    if path != "" && action != "" {

        if action == "-r"{
            fmt.Println(datadriver.ReadData(path))
        }

        if action == "-w" {
            data = strings.Join(os.Args[3:], " ")
            if data != "" {
                datadriver.WriteData(path, data)
            } else {
                fmt.Println("No data to write")
            }
        }

    } else {
        fmt.Println("No path to file provided!")
    }
}
