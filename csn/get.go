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
    --qual [quality]    - max quality to print out (default to 'all')
                          either '32', '128', '320', '500', 'lossless' or 'all'
    --help              - show this help

Examples:
    %[1]s get --qual lossless https://chiasenhac.vn/nghe-album/...
    %[1]s get --qual all https://vn.chiasenhac.vn/mp3/vietnam/v-pop/...
`, os.Args[0])
	os.Exit(0)
}

func doGet(qual string, target string) {
	info, err := csn.GetInfoUrl(target)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Also see https://github.com/l4rzy/csn-ng/issues/1")
	}
	info.PrintLinks(csn.MUSIC_QUAL_ALL)
}
