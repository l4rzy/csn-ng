/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package main

import (
	"fmt"
	csn "github.com/l4rzy/csn-ng"
	"os"
)

func helpSearch() {
	fmt.Printf(`csn search - search for data on chiasenhac.vn
Synopsis:
    %[1]s search [flags] [input]

Available flags:
    --music         - search for music (I mean, songs)
    --video         - search for videos
    --artist        - search for videos
    --album         - search for albums
    --link          - also get download links for each result
    --limit [int]   - limit search results for each entry (5 by default)
    --help          - show this help

Examples:
    %[1]s search --limit 10 --music --link "what I've done"
    %[1]s search -album -music love
`, os.Args[0])
	os.Exit(0)
}

func doSearch(opt int, keyword string, limit int, link bool) {
	result, err := csn.SearchNew(opt, keyword, limit)
	if err != nil {
		fmt.Printf("Could not get data: %v\n", err)
		os.Exit(-1)
	}

	for _, r := range result {
		r.Print()
		if link {
			info, _ := r.GetInfo()
			info.PrintLinks(csn.MUSIC_QUAL_ALL)
		}
		fmt.Println("")
	}
}
