/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package main

import (
	"flag"
	"fmt"
	csn "github.com/l4rzy/csn-ng"
	"os"
)

func showAbout() {
	fmt.Printf(`csn - do things with chiasenhac.vn
version: %d.%d.%d
`, csn.VER_MAJOR, csn.VER_MINOR, csn.VER_PATCH)
	os.Exit(0)
}

func showHelp() {
	fmt.Printf(`csn - do things with chiasenhac.vn
Example:
  %[1]s -limit 10 -album "linkin"
  %[1]s -limit 5 -artist "tuấn"
  %[1]s -album -music "love"

type '%[1]s --help' for help
`, os.Args[0])
	os.Exit(0)
}

func main() {
	var (
		s_about  = flag.Bool("version", false, "show program info")
		s_music  = flag.Bool("music", false, "search for music")
		s_video  = flag.Bool("video", false, "search for video")
		s_artist = flag.Bool("artist", false, "search for artist")
		s_album  = flag.Bool("album", false, "search for album")
		s_link   = flag.Bool("link", false, "also get the download links of music")
		s_limit  = flag.Int("limit", 5, "limit search result on each category")
	)

	flag.Parse()

	if *s_about {
		showAbout()
	}

	// construct option
	var opt = 0
	if *s_music {
		opt |= csn.OP_MUSIC
	}
	if *s_video {
		opt |= csn.OP_VIDEO
	}
	if *s_artist {
		opt |= csn.OP_ARTIST
	}
	if *s_album {
		opt |= csn.OP_ALBUM
	}

	if len(flag.Args()) == 0 {
		showHelp()
		os.Exit(-1)
	}

	if opt == 0 {
		showHelp()
		os.Exit(-1)
	}

	var keyword = flag.Args()[0]
	result, err := csn.SearchNew(opt, keyword, *s_limit)
	if err != nil {
		fmt.Printf("Could not get data: %v\n", err)
		os.Exit(-1)
	}

	for _, r := range result {
		r.Print()
		if *s_link {
			info, _ := r.GetInfo()
			fmt.Printf("\t[128] %v\n", info.FileURL)
			fmt.Printf("\t[320] %v\n", info.File320URL)
			fmt.Printf("\t[500] %v\n", info.FileM4AURL)
			fmt.Printf("\t[Lossless] %v\n", info.FileLosslessURL)
		}
		fmt.Println("")
	}
}
