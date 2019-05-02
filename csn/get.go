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
	var (
		qopt int = 0
	)
	switch (qual) {
	case "32":
		qopt |= csn.MUSIC_QUAL_32
	case "128":
		qopt |= csn.MUSIC_QUAL_128
	case "320":
		qopt |= csn.MUSIC_QUAL_320
	case "500":
		qopt |= csn.MUSIC_QUAL_500
	case "lossless":
		qopt |= csn.MUSIC_QUAL_1000
	case "all":
		qopt |= csn.MUSIC_QUAL_ALL
	}

	info, err := csn.GetInfoUrl(target)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Also see https://github.com/l4rzy/csn-ng/issues/1")
	}
	info.PrintLinks(false, qopt)
}
