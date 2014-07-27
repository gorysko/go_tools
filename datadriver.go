package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "strings"
    "os"
    // "io"
    "compress/zlib"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read_file(path string) string {
    dat, err := ioutil.ReadFile(path)
    check(err)
    buffer := uncompress_data(dat)
    return string(buffer.String())

}


func write_data(path string, data string) {
    write_data := []byte(data)
    compressed := compress_data(write_data)
    err := ioutil.WriteFile(path, compressed.Bytes(), 0644)
    check(err)
}


func compress_data(data []byte) bytes.Buffer {
    var write_bytes bytes.Buffer

    writer := zlib.NewWriter(&write_bytes)
    writer.Write(data)
    writer.Close()
    return write_bytes
}

func uncompress_data(data []byte) bytes.Buffer {
    var buffer bytes.Buffer

    buff := []byte(data)

    reader := bytes.NewReader(buff)
    read, err := zlib.NewReader(reader)
    check(err)
    buffer.ReadFrom(read)
    return buffer
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
