# Description
* A shitty but working API and CLI tools for searching (and maybe more) on chiasenhac.vn
* Please, just use it, don't look at the code

## Build

```sh
$ go get 'github.com/l4rzy/csn-ng/csn'
```

## Usage
* As a library
```go
package main

import (
    "fmt"
    csn "github.com/l4rzy/csn-ng"
    "os"
)

func main() {
    keyword := os.Args[1]
    result, err := csn.SearchNew(csn.KIND_MUSIC, keyword, 10)
    if err != nil {
        fmt.Printf("Could not get data: %v\n", err)
        os.Exit(-1)
    }

    for _, r := range result {
        r.Print()
        info, _ := r.GetInfo()
        info.PrintLinks(true, csn.MUSIC_QUAL_ALL)
    }
}

```

* As a commandline tool

Make sure that `$GOPATH/bin` is in your `$PATH`
```sh
$ csn search -link -limit 3 -music "what I've done"
```
The result should look like this
![test](test.png)

## License
MIT
