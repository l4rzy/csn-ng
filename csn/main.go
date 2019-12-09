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
	"runtime"
)

func showAbout() {
	fmt.Printf(`csn - do things with chiasenhac.vn
runtime: %s
version: %v.%v.%v
bug report: https://github.com/l4rzy/csn-ng/issues
`, runtime.Version(), csn.VER_MAJOR, csn.VER_MINOR, csn.VER_PATCH)
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
		sub_search = flag.NewFlagSet("search", flag.ExitOnError)
		sub_get    = flag.NewFlagSet("get", flag.ExitOnError)
		p_about    = flag.Bool("version", false, "show program info")
		p_help     = flag.Bool("help", false, "show help")
	)

	// flags for search
	var (
		s_help   = sub_search.Bool("help", false, "show help for search")
		s_music  = sub_search.Bool("music", false, "search for music")
		s_video  = sub_search.Bool("video", false, "search for video")
		s_artist = sub_search.Bool("artist", false, "search for artist")
		s_album  = sub_search.Bool("album", false, "search for album")
		s_link   = sub_search.Bool("link", false, "also get the download links of music")
		s_limit  = sub_search.Int("limit", 5, "limit search result on each category")
	)

	// flags for get
	var (
		g_help    = sub_get.Bool("help", false, "show help for get")
		g_quality = sub_get.String("qual", "320", "max quality to get")
		g_batch   = sub_get.Bool("file", false, "read links from file")
	)

	// sub functions
	var cmdSearch = func() {
		if *s_help {
			helpSearch()
		}

		// construct option
		var opt = 0
		if *s_music {
			opt |= csn.KIND_MUSIC
		}
		if *s_video {
			opt |= csn.KIND_VIDEO
		}
		if *s_artist {
			opt |= csn.KIND_ARTIST
		}
		if *s_album {
			opt |= csn.KIND_ALBUM
		}

		if opt == 0 {
			opt |= csn.KIND_MUSIC
		}

		if len(sub_search.Args()) == 0 {
			helpSearch()
		}

		var keyword = sub_search.Args()[0]
		doSearch(opt, keyword, *s_limit, *s_link)
	}

	var cmdGet = func() {
		if *g_help {
			helpGet()
		}
		if *g_batch {
			fmt.Println("Currently not available")
			os.Exit(0)
		}
		if len(sub_get.Args()) == 0 {
			helpGet()
		}

		var target = sub_get.Args()[0]
		doGet(*g_quality, target)
	}

	// override usages
	flag.Usage = showHelp
	sub_search.Usage = helpSearch
	sub_get.Usage = helpGet

	flag.Parse()
	if *p_about {
		showAbout()
	}

	if *p_help {
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
		sub_search.Parse(os.Args[2:])
		cmdSearch()
	case "get", "g":
		if len(os.Args) == 2 {
			helpGet()
		}
		sub_get.Parse(os.Args[2:])
		cmdGet()
	default:
		showHelp()
	}
}
