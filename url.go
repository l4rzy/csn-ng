/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

import (
	"strings"
    "errors"
)

func withSchemePrefix(url string) (CSNUrlInfo, error) {
    var ret CSNUrlInfo
    ret.Url = url

    toks := strings.Split(url, "/")

    switch toks[3] {
    case "mp3":
        ret.Kind = KIND_MUSIC
        ret.UrlName = toks[6][:strings.LastIndex(toks[6], "~")]
        ret.Category = toks[4]
    case "nghe-album":
        ret.Kind = KIND_ALBUM
        ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
    case "hd":
        ret.Kind = KIND_VIDEO
        ret.UrlName = toks[6][:strings.LastIndex(toks[6], "~")]
    }

    return ret, nil
}

func extractUrlInfo(url string) (CSNUrlInfo, error) {
	var ret CSNUrlInfo
    ret.Url = "https://" + url
    if strings.Contains(url, "old.chiasenhac.vn") {
        err := errors.New("doesn't support old site")
        return ret, err
    }
    if strings.HasPrefix(url, "https://") {
        return withSchemePrefix(url)
    }
	toks := strings.Split(url, "/")

    switch toks[1] {
    case "mp3":
        ret.Kind = KIND_MUSIC
        ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
        ret.Category = toks[2]
    case "nghe-album":
        ret.Kind = KIND_ALBUM
        ret.UrlName = toks[2][:strings.LastIndex(toks[2], "~")]
    case "hd":
        ret.Kind = KIND_VIDEO
        ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
    case "playlist":
        ret.Kind = KIND_PLAYLIST
        ret.UrlName = toks[2][:strings.LastIndex(toks[2], "~")]
    }
	return ret, nil
}

func GetInfoUrl(url string) (CSNMusicInfo, error) {
    var (
        ret CSNMusicInfo
        id interface{}
    )
    uinfo, err := extractUrlInfo(url)
    if err != nil {
        return ret, err
    }

    sret, err := SearchNew(KIND_MUSIC|KIND_ALBUM|KIND_VIDEO, uinfo.UrlName, 10)
    for _, i := range sret {
        if (i.GetLink() == uinfo.Url) {
            id = i.GetID()
        }
    }

    return getInfo(id)
}
