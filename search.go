/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (m CSNMusicSearch) Print() {
	fmt.Printf("%v\n", m.MusicLink)
}

func (v CSNVideoSearch) Print() {
	fmt.Printf("%v\n", v.VideoLink)
}

func (a CSNArtistSearch) Print() {
	fmt.Printf("%s%v\n", CSN_HOME, a.ArtistLink)
}

func (a CSNAlbumSearch) Print() {
	fmt.Printf("%s%v\n", CSN_HOME, a.AlbumLink)
}

func buildSearchUrl(keyword string, limit int) (string, error) {
	if limit <= 0 || keyword == "" {
		err := errors.New("Wrong search input")
		return "", err
	}
	return fmt.Sprintf(SEARCH_FMT, keyword, limit), nil
}

func Search(opt int, keyword string, limit int) ([]CSNObjectSearch, error) {
	surl, err := buildSearchUrl(keyword, limit)
	if err != nil {
		return nil, err
	}

	// get data
	resp, err := csn_client.Get(surl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Responsed status code is not ok")
		return nil, err
	}

	// read to body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse json
	var result CSNSearchResp
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	var data = result[0]

	// return result as user asked
	var ret []CSNObjectSearch
	if opt&OP_MUSIC != 0 {
		for i := 0; i < len(data.Music.Data); i++ {
			ret = append(ret, data.Music.Data[i])
		}
	}
	if opt&OP_VIDEO != 0 {
		for i := 0; i < len(data.Video.Data); i++ {
			ret = append(ret, data.Video.Data[i])
		}
	}
	if opt&OP_ARTIST != 0 {
		for i := 0; i < len(data.Artist.Data); i++ {
			ret = append(ret, data.Artist.Data[i])
		}
	}
	if opt&OP_ALBUM != 0 {
		for i := 0; i < len(data.Album.Data); i++ {
			ret = append(ret, data.Album.Data[i])
		}
	}
	return ret, nil
}
