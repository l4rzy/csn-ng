/*
 * Copyright (C) 2023 l4rzy
 * MIT License
 */

package csn

import (
	"net/http"
	"time"
)

const (
	VerMajor  = 0
	VerMinor = 1
	VerPatch = 4
)

var (
	csnClient *http.Client
)

// search options
const (
	KindMusic   = 1 << 1
	KindVideo  = 1 << 2
	KindArtist    = 1 << 3
	KindAlbum    = 1 << 4
	KindPlaylist = 1 << 5
)

// music quality
const (
	MusicQual32    = 1 << 1
	MusicQual128   = 1 << 2
	MusicQual320    = 1 << 3
	MusicQual500   = 1 << 4
	MusicQual1000 = 1 << 5
	MusicQualAll  = 62
)

const (
    Home         = "https://chiasenhac.vn/"
	SearchNewFmt = "https://chiasenhac.vn/search/real?q=%v&type=json&rows=%v&view_all=true"
	SearchFmt    = "http://search.chiasenhac.vn/api/search.php?s=%v&code=csn22052018&per_page=%v"
	MusicInfoFmt = "http://old.chiasenhac.vn/api/listen.php?code=csn22052018&return=json&m=%v"
)

func init() {
	csnClient = &http.Client{Timeout: 10 * time.Second}
}
