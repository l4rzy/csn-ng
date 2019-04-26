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
	CSN_HOME   = "https://chiasenhac.vn"
	SEARCH_FMT = "https://chiasenhac.vn/search/real?q=%s&type=json&rows=%d&view_all=true"
)

func init() {
	csn_client = &http.Client{Timeout: 10 * time.Second}
}
