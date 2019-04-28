/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package main

import (
	"fmt"
	"os"
	csn "github.com/l4rzy/csn-ng"
)

func helpGet() {
	fmt.Printf(`csn get - get download links for song, album, playlist
Synopsis:
    %[1]s get [flags] [input]

Available flags:
    --qual [quality]    - max quality to get (default to 320)
                          either '32', '128', '320', '500' or 'lossless'
    --help              - show this help

Examples:
    %[1]s get --qual lossless https://chiasenhac.vn/nghe-album/...
    %[1]s get --qual all https://vn.chiasenhac.vn/mp3/vietnam/v-pop/...
`, os.Args[0])
	os.Exit(0)
}

func doGet(qual string, target string) {
	csn.ExtractUrlInfo(target)
}
