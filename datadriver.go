package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "os"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read_file(path string) string {
    dat, err := ioutil.ReadFile(path)
    check(err)
    return string(dat)

}

func write_data(path string, data string) {
    write_data := []byte(data)
    err := ioutil.WriteFile(path, write_data, 0644)
    check(err)
}

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
            fmt.Println(read_file(path))
        }

        if action == "-w" {
            data = strings.Join(os.Args[3:], " ")
            if data != "" {
                write_data(path, data)
            } else {
                fmt.Println("No data to write")
            }
        }

    } else {
        fmt.Println("No path to file provided!")
    }
}
