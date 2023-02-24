/*
 * Copyright (C) 2023 l4rzy
 * MIT License
 */

package main

import (
	"flag"
	"fmt"
	csn "github.com/l4rzy/csn-ng"
	"os"
	"runtime"
)

func showAbout() {
	fmt.Printf(`csn - do things with chiasenhac.vn
runtime: %s
version: %v.%v.%v
bug report: https://github.com/l4rzy/csn-ng/issues
`, runtime.Version(), csn.VerMajor, csn.VerMinor, csn.VerPatch)
	os.Exit(0)
}

func showHelp() {
	fmt.Printf(`csn - do things with chiasenhac.vn
Synopsis:
    %[1]s [subcommand] [flags] [input]
    %[1]s [flags]

Available subcommands:
    search (s)     - search data on chiasenhac.vn
    get    (g)     - get download links for songs, videos, albums, playlists

Available top-level flags:
    --help         - show this help
    --version      - show program version

Examples:
    %[1]s search -limit 10 -album "linkin"
    %[1]s s -album -music "love"
    %[1]s get -qual lossless https://chiasenhac.vn/nghe-album/...
    %[1]s g https://chiasenhac.vn/nghe-album/...
`, os.Args[0])
	os.Exit(0)
}

func main() {

	// top level flags and subcommands
	var (
		subSearch = flag.NewFlagSet("search", flag.ExitOnError)
		subGet = flag.NewFlagSet("get", flag.ExitOnError)
		pAbout = flag.Bool("version", false, "show program info")
		pHelp  = flag.Bool("help", false, "show help")
	)

	// flags for search
	var (
		sHelp   = subSearch.Bool("help", false, "show help for search")
		sMusic   = subSearch.Bool("music", false, "search for music")
		sVideo  = subSearch.Bool("video", false, "search for video")
		sArtist = subSearch.Bool("artist", false, "search for artist")
		sAlbum  = subSearch.Bool("album", false, "search for album")
		sLink  = subSearch.Bool("link", false, "also get the download links of music")
		sLimit = subSearch.Int("limit", 5, "limit search result on each category")
	)

	// flags for get
	var (
		gHelp    = subGet.Bool("help", false, "show help for get")
		gQuality = subGet.String("qual", "320", "max quality to get")
		gBatch   = subGet.Bool("file", false, "read links from file")
	)

	// sub functions
	var cmdSearch = func() {
		if *sHelp {
			helpSearch()
		}

		// construct option
		var opt = 0
		if *sMusic {
			opt |= csn.KindMusic
		}
		if *sVideo {
			opt |= csn.KindVideo
		}
		if *sArtist {
			opt |= csn.KindArtist
		}
		if *sAlbum {
			opt |= csn.KindAlbum
		}

		if opt == 0 {
			opt |= csn.KindMusic
		}

		if len(subSearch.Args()) == 0 {
			helpSearch()
		}

		var keyword = subSearch.Args()[0]
		doSearch(opt, keyword, *sLimit, *sLink)
	}

	var cmdGet = func() {
		if *gHelp {
			helpGet()
		}
		if *gBatch {
			fmt.Println("Currently not available")
			os.Exit(0)
		}
		if len(subGet.Args()) == 0 {
			helpGet()
		}

		var target = subGet.Args()[0]
		doGet(*gQuality, target)
	}

	// override usages
	flag.Usage = showHelp
	subSearch.Usage = helpSearch
	subGet.Usage = helpGet

	flag.Parse()
	if *pAbout {
		showAbout()
	}

	if *pHelp {
		showHelp()
	}

	// subcommands
	if len(os.Args) == 1 {
		showHelp()
	}

	switch os.Args[1] {
	case "search", "s":
		if len(os.Args) == 2 {
			helpSearch()
		}
		subSearch.Parse(os.Args[2:])
		cmdSearch()
	case "get", "g":
		if len(os.Args) == 2 {
			helpGet()
		}
		subGet.Parse(os.Args[2:])
		cmdGet()
	default:
		showHelp()
	}
}
