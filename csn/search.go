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
    %[1]s s [flags] [input]

Available flags:
    --music         - search for music (songs, default mode)
    --video         - search for videos
    --artist        - search for artists
    --album         - search for albums
    --link          - also get download links for each result, in case of songs and videos
    --limit [int]   - limit search results for each entry (5 by default)
    --help          - show this help

Examples:
    %[1]s s --link "in the end"
    %[1]s search --limit 10 --music --link "what I've done"
    %[1]s search -album -music love
`, os.Args[0])
	os.Exit(0)
}

func PrintLinks(mi csn.CSNMusicInfo, prefix bool, opt int) {
	if prefix == true {
		if opt&csn.MUSIC_QUAL_32 != 0 &&
			mi.File32URL != "" {
			fmt.Printf("[ 32] %v\n", mi.File32URL)
		}
		if opt&csn.MUSIC_QUAL_128 != 0 &&
			mi.FileURL != "" {
			fmt.Printf("[128] %v\n", mi.FileURL)
		}
		if opt&csn.MUSIC_QUAL_320 != 0 &&
			mi.File320URL != "" {
			fmt.Printf("[320] %v\n", mi.File320URL)
		}
		if opt&csn.MUSIC_QUAL_500 != 0 &&
			mi.FileM4AURL != "" {
			fmt.Printf("[500] %v\n", mi.FileM4AURL)
		}
		if opt&csn.MUSIC_QUAL_1000 != 0 &&
			mi.FileLosslessURL != "" {
			fmt.Printf("[LLs] %v\n", mi.FileLosslessURL)
		}
	} else {
		if opt&csn.MUSIC_QUAL_32 != 0 &&
			mi.File32URL != "" {
			fmt.Printf("%v\n", mi.File32URL)
		}
		if opt&csn.MUSIC_QUAL_128 != 0 &&
			mi.FileURL != "" {
			fmt.Printf("%v\n", mi.FileURL)
		}
		if opt&csn.MUSIC_QUAL_320 != 0 &&
			mi.File320URL != "" {
			fmt.Printf("%v\n", mi.File320URL)
		}
		if opt&csn.MUSIC_QUAL_500 != 0 &&
			mi.FileM4AURL != "" {
			fmt.Printf("%v\n", mi.FileM4AURL)
		}
		if opt&csn.MUSIC_QUAL_1000 != 0 &&
			mi.FileLosslessURL != "" {
			fmt.Printf("%v\n", mi.FileLosslessURL)
		}
	}
}

func doSearch(opt int, keyword string, limit int, link bool) {
	result, err := csn.SearchNew(opt, keyword, limit)
	if err != nil {
		fmt.Printf("Could not get data: %v\n", err)
		os.Exit(-1)
	}

	//fmt.Printf("%#v\n", result)

	for _, r := range result {
		r.Print()
		if link {
			info, _ := r.GetInfo()
			PrintLinks(info, true, csn.MUSIC_QUAL_ALL)
		}
		fmt.Println("")
	}
}
