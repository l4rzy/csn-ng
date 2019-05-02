/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

import (
	"net/http"
	"time"
)

const (
	VER_MAJOR = 0
	VER_MINOR = 1
	VER_PATCH = 1
)

var (
	csn_client *http.Client
	err        error
)

// search options
const (
	KIND_MUSIC  = 1 << 1
	KIND_VIDEO  = 1 << 2
	KIND_ARTIST = 1 << 3
	KIND_ALBUM  = 1 << 4
	KIND_PLAYLIST = 1 << 5
)

// music quality
const (
	MUSIC_QUAL_32   = 1 << 1
	MUSIC_QUAL_128  = 1 << 2
	MUSIC_QUAL_320  = 1 << 3
	MUSIC_QUAL_500  = 1 << 4
	MUSIC_QUAL_1000 = 1 << 5
	MUSIC_QUAL_ALL  = 62
)

const (
	CSN_HOME       = "https://chiasenhac.vn"
	SEARCH_NEW_FMT = "https://chiasenhac.vn/search/real?q=%v&type=json&rows=%v&view_all=true"
	SEARCH_FMT     = "http://search.chiasenhac.vn/api/search.php?s=%v&code=csn22052018&per_page=%v"
	MUSIC_INFO_FMT = "http://old.chiasenhac.vn/api/listen.php?code=csn22052018&return=json&m=%v"
)

func init() {
	csn_client = &http.Client{Timeout: 10 * time.Second}
}
