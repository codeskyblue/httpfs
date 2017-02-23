# httpfs
[![GoDoc](https://godoc.org/github.com/codeskyblue/httpfs?status.svg)](https://godoc.org/github.com/codeskyblue/httpfs)

Golang library for open http file like open local file

## Example
Usage in unzip http file

```go
package main

import (
    "github.com/codeskyblue/httpfs"
    "archive/zip"
    "fmt"
    "log"
)

func main(){
    file, err := Open("http://some-host/file.zip")
	if err != nil {
		log.Fatal(err)
	}
    zrd, err := zip.NewReader(file, file.Size())
    if err != nil {
		log.Fatal(err)
	}
    for _, f := range zrd.File {
        fmt.Printf("F-Name: %s\n", f.Name)
    }
}
```

## LICENSE
[MIT](LICENSE)