package main

import (
        "fmt"
        "os"
        "datadriver"
)

type ApacheRow struct {
    ip_adress string
    domain string
    user string
    datetime string
    method string
    uri string
    proto string
    status int
    bytes string
    referrer string
    agent string
    cookie string
}

const  format string = "%s %s %s [%[^]]] \"%s %s %[^\"]\" %d %s"

func apache_parse(line string) *ApacheRow {
    result := new(ApacheRow)
    fmt.Sscanf(line, format, &result.ip_adress, &result.domain,
                &result.user, &result.datetime, &result.method, &result.uri,
                &result.proto, &result.status, &result.bytes, &result.referrer,
                &result.agent, &result.cookie)
    if result.status == 200 || result.status == 302 || result.status == 304 {
        return result
    }
    return nil
}

func main() {
    if len(os.Args) == 0 {
        fmt.Println("No files was provided")
    }
    data := datadriver.ReadData(os.Args[1])
    if data != "" {
        fmt.Println(data)
    }

    fmt.Println("No data")
}
