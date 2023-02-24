/*
 * Copyright (C) 2023 l4rzy
 * MIT License
 */

package csn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

func (m MusicSearch) Print() {
	fmt.Printf("[%v] %v (%v)\n", m.MusicID, m.MusicTitle, m.MusicArtist)
}

func (m MusicSearchNew) Print() {
	fmt.Printf("[%v] %v (%v)\n", m.MusicID, m.MusicTitle, m.MusicArtist)
}

func (v VideoSearch) Print() {
	fmt.Printf("%v\n", v.VideoLink)
}

func (a ArtistSearch) Print() {
	fmt.Printf("%s%v\n", Home, a.ArtistLink)
}

func (a AlbumSearch) Print() {
	fmt.Printf("%s%v\n", Home, a.AlbumLink)
}

func SearchNew(opt int, keyword string, limit int) ([]ObjectSearch, error) {
	if limit <= 0 || keyword == "" {
		err := errors.New("wrong search input")
		return nil, err
	}
	surl := fmt.Sprintf(SearchNewFmt, url.PathEscape(keyword), limit)

	// get data
	resp, err := csnClient.Get(surl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("responsed status code is not ok")
		return nil, err
	}

	// read to body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse json
	var result CSNSearchNewResp
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	var data = result[0]

	// return result as user asked
	var ret []ObjectSearch

	if opt&KindMusic != 0 {
		for i := 0; i < len(data.Music.Data); i++ {
			ret = append(ret, data.Music.Data[i])
		}
	}
	if opt&KindVideo != 0 {
		for i := 0; i < len(data.Video.Data); i++ {
			ret = append(ret, data.Video.Data[i])
		}
	}
	if opt&KindArtist != 0 {
		for i := 0; i < len(data.Artist.Data); i++ {
			ret = append(ret, data.Artist.Data[i])
		}
	}
	if opt&KindAlbum != 0 {
		for i := 0; i < len(data.Album.Data); i++ {
			ret = append(ret, data.Album.Data[i])
		}
	}

	return ret, nil
}

func Search(keyword string, limit int) ([]ObjectSearch, error) {
	if limit <= 0 || keyword == "" {
		err := errors.New("wrong search input")
		return nil, err
	}
	surl := fmt.Sprintf(SearchFmt, url.PathEscape(keyword), limit)

	// get data
	resp, err := csnClient.Get(surl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("responsed status code is not ok")
		return nil, err
	}

	// read to body
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// parse json
	var result CSNSearchResp
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	var ret []ObjectSearch
	for i := 0; i < len(result.MusicList); i++ {
		ret = append(ret, result.MusicList[i])
	}

	return ret, nil
}
