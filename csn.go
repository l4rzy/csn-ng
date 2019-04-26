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

const (
	OP_MUSIC  = 1 << 1
	OP_VIDEO  = 1 << 2
	OP_ARTIST = 1 << 3
	OP_ALBUM  = 1 << 4
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
