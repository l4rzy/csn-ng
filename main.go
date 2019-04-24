package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VER_MAJ = 0
	VER_MIN = 1
)

func showAbout() {
	fmt.Printf(`csn - do things with chiasenhac.vn
version: %d.%d
`, VER_MAJ, VER_MIN)
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
		s_limit  = flag.Int("limit", 5, "limit search result on each category")
	)

	flag.Parse()

	if *s_about {
		showAbout()
	}

	// construct option
	var opt = 0
	if *s_music {
		opt |= OP_MUSIC
	}
	if *s_video {
		opt |= OP_VIDEO
	}
	if *s_artist {
		opt |= OP_ARTIST
	}
	if *s_album {
		opt |= OP_ALBUM
	}

	if len(flag.Args()) == 0 {
		showHelp()
		os.Exit(-1)
	}

	var keyword = flag.Args()[0]
	result, err := search(opt, keyword, *s_limit)
	if err != nil {
		fmt.Printf("Could not get data: %v\n", err)
		os.Exit(-1)
	}

	for i, r := range result {
		fmt.Printf("[%d] ", i+1)
		r.print()
	}

}
