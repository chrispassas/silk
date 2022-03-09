[![GoDoc](https://godoc.org/github.com/chrispassas/silk?status.svg)](https://godoc.org/github.com/chrispassas/silk)
[![Go Report Card](https://goreportcard.com/badge/github.com/chrispassas/silk)](https://goreportcard.com/report/github.com/chrispassas/silk)


# silk flows
https://tools.netsa.cert.org/silk/docs.html

## Description
This package makes it easy to read common silk files without using C Go.

## Go Doc
https://godoc.org/github.com/chrispassas/silk

## What is silk
Source https://tools.netsa.cert.org/silk/faq.html#what-silk 
>SiLK is a suite of network traffic collection and analysis tools developed and maintained by the CERT Network Situational >Awareness Team (CERT NetSA) at Carnegie Mellon University to facilitate security analysis of large networks. The SiLK tool >suite supports the efficient collection, storage, and analysis of network flow data, enabling network security analysts to >rapidly query large historical traffic data sets.

## Supported File Formats
| Record Size   | Record Version           | Compression   | Supported          |
| ------------- | -------------            | ------------- | -------------      |
| 88            | RWIPV6ROUTING VERSION 1  | None (0)      | :white_check_mark: |
| 88            | RWIPV6ROUTING VERSION 1  | Zlib (1)      | :white_check_mark: |
| 88            | RWIPV6ROUTING VERSION 1  | Lzo (2)       | :white_check_mark: |
| 88            | RWIPV6ROUTING VERSION 1  | Snappy (3)    | :white_check_mark: |
| 68            | RWIPV6 VERSION 1         | None (0)      | :white_check_mark: |
| 68            | RWIPV6 VERSION 1         | Zlib (1)      | :white_check_mark: |
| 68            | RWIPV6 VERSION 1         | Lzo (2)       | :white_check_mark: |
| 68            | RWIPV6 VERSION 1         | Snappy (3)    | :white_check_mark: |
| 56            | RWIPV6 VERSION 2         | None (0)      | :white_check_mark: |
| 56            | RWIPV6 VERSION 2         | Zlib (1)      | :white_check_mark: |
| 56            | RWIPV6 VERSION 2         | Lzo (2)       | :white_check_mark: |
| 56            | RWIPV6 VERSION 2         | Snappy (3)    | :white_check_mark: |
| 52            | RWGENERIC VERSION 5      | None (0)      | :white_check_mark: |
| 52            | RWGENERIC VERSION 5      | Zlib (1)      | :white_check_mark: |
| 52            | RWGENERIC VERSION 5      | Lzo (2)       | :white_check_mark: |
| 52            | RWGENERIC VERSION 5      | Snappy (3)    | :white_check_mark: |

## Example

### Parse Whole File
```go
package main

import (
    "fmt"
    "log"
    "github.com/chrispassas/silk"
)

func main() {

    var testFile = "testdata/FT_RWIPV6-v2-c0-L.dat"
    var err error
    var sf silk.File

    if sf, err = silk.OpenFile(testFile); err != nil {
        log.Fatalf("OpenFile() error:%s", err)
    }

    log.Printf("Compression:%d", sf.Header.Compression)
    log.Printf("FileFlags:%d", sf.Header.FileFlags)
    log.Printf("FileVersion:%d", sf.Header.FileVersion)
    log.Printf("HeaderLength:%d", sf.Header.HeaderLength)
    log.Printf("MagicNumber:%x", sf.Header.MagicNumber)
    log.Printf("RecordFormat:%d", sf.Header.RecordFormat)
    log.Printf("RecordSize:%d", sf.Header.RecordSize)
    log.Printf("RecordVersion:%d", sf.Header.RecordVersion)
    log.Printf("SilkVersion:%d", sf.Header.SilkVersion)

    log.Printf("File record count:%d\n", len(sf.Flows))

    fmt.Printf("start_time_ms,src_ip,dst_ip,src_port,dst_port\n")
    for _, flow := range sf.Flows {
        fmt.Printf("%d,%s,%s,%d,%d\n",
            flow.StartTimeMS,
            flow.SrcIP.String(),
            flow.DstIP.String(),
            flow.SrcPort,
            flow.DstPort,
        )
        //Etc... for other silk.Flow values
    }
}
```

### Channel Based Parsing
```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/chrispassas/silk"
)

func main() {
    var testFile = "testdata/FT_RWIPV6-v2-c0-L.dat"
    var err error
    
    flows := make([]silk.Flow, 0, 245340)
    receiver := silk.NewChannelFlowReceiver(0)

    reader, err := os.Open(testFile)
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()

    go func() {
        if err = silk.Parse(reader, receiver); err != nil {
            log.Fatal(err)
        }
    }()

    for flow := range receiver.Read() {
        /*
            Pulling all data into an in memory array. That really isn't the point of the channel based
            parser. You would want to stream it somewhere else to keep memory usage low. This is for example
            purposes only.
        */
        flows = append(flows, flow)
    }

    log.Printf("Compression:%d", receiver.Header.Compression)
    log.Printf("FileFlags:%d", receiver.Header.FileFlags)
    log.Printf("FileVersion:%d", receiver.Header.FileVersion)
    log.Printf("HeaderLength:%d", receiver.Header.HeaderLength)
    log.Printf("MagicNumber:%x", receiver.Header.MagicNumber)
    log.Printf("RecordFormat:%d", receiver.Header.RecordFormat)
    log.Printf("RecordSize:%d", receiver.Header.RecordSize)
    log.Printf("RecordVersion:%d", receiver.Header.RecordVersion)
    log.Printf("SilkVersion:%d", receiver.Header.SilkVersion)

    log.Printf("File record count:%d\n", len(flows))

    fmt.Printf("start_time_ms,src_ip,dst_ip,src_port,dst_port\n")
    for _, flow := range flows {
        fmt.Printf("%d,%s,%s,%d,%d\n",
            flow.StartTimeMS,
            flow.SrcIP.String(),
            flow.DstIP.String(),
            flow.SrcPort,
            flow.DstPort,
        )
        //Etc... for other silk.Flow values
    }
}
```
