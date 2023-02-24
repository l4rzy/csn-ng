/*
 * Copyright (C) 2023 l4rzy
 * MIT License
 */

package csn

import (
	"errors"
	"fmt"
	"strings"
)

func withSchemePrefix(url string) (UrlInfo, error) {
	var ret UrlInfo
	ret.Url = url

	toks := strings.Split(url, "/")

	switch toks[3] {
	case "mp3":
		ret.Kind = KindMusic
		ret.UrlName = toks[6][:strings.LastIndex(toks[6], "~")]
		ret.Category = toks[4]
	case "nghe-album":
		ret.Kind = KindAlbum
		ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
	case "hd":
		ret.Kind = KindVideo
		ret.UrlName = toks[6][:strings.LastIndex(toks[6], "~")]
	}

	return ret, nil
}

func extractUrlInfo(url string) (UrlInfo, error) {
	var ret UrlInfo
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
		ret.Kind = KindMusic
		ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
		ret.Category = toks[2]
	case "nghe-album":
		ret.Kind = KindAlbum
		ret.UrlName = toks[2][:strings.LastIndex(toks[2], "~")]
	case "hd":
		ret.Kind = KindVideo
		ret.UrlName = toks[4][:strings.LastIndex(toks[4], "~")]
	case "playlist":
		ret.Kind = KindPlaylist
		ret.UrlName = toks[2][:strings.LastIndex(toks[2], "~")]
	}
	return ret, nil
}

func GetInfoUrl(url string) (MusicInfo, error) {
	var (
		ret MusicInfo
		id  interface{}
	)
	uinfo, err := extractUrlInfo(url)
	if err != nil {
		return ret, err
	}
	if uinfo.Kind == KindPlaylist {
		err := errors.New("Currently doesn't support getting playlists")
		return ret, err
	}

	sret, err := SearchNew(KindMusic|KindAlbum|KindVideo, uinfo.UrlName, 10)
	for _, i := range sret {
		if i.GetLink() == uinfo.Url {
			id = i.GetID()
		}
	}

	switch uinfo.Kind {
	case KindPlaylist:
		err := errors.New("Currently doesn't support getting playlists")
		return ret, err

	case KindVideo, KindMusic:
		sret, err := SearchNew(KindMusic|KindVideo, uinfo.UrlName, 10)
		if err != nil {
			return ret, err
		}
		for _, i := range sret {
			if i.GetLink() == uinfo.Url {
				id = i.GetID()
			}
		}

		return getInfo(id)
	case KindAlbum:
		sret, err := SearchNew(KindAlbum, uinfo.UrlName, 20)
		//fmt.Printf("%#v\n", sret)
		if err != nil {
			return ret, err
		}
		for _, i := range sret {
			if i.GetLink() == uinfo.Url {
				fmt.Printf("Found %#v\n", i)
			}
		}
	}

	return ret, nil
}
