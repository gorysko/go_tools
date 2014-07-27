package datadriver

import (
    "bytes"
    "io/ioutil"
    "compress/zlib"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadData(path string) string {
    dat, err := ioutil.ReadFile(path)
    check(err)
    buffer := uncompress_data(dat)
    return string(buffer.String())

}

func WriteData(path string, data string) {
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

