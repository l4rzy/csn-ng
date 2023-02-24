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
)

func getInfo(id interface{}) (MusicInfo, error) {
	var ret MusicInfoResp
	surl := fmt.Sprintf(MusicInfoFmt, id)

	// get data
	resp, err := csnClient.Get(surl)
	if err != nil {
		return ret.MusicInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("responsed status code is not ok")
		return ret.MusicInfo, err
	}

	// read to body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ret.MusicInfo, err
	}

	// parse json
	err = json.Unmarshal(body, &ret)

	if err != nil {
		return ret.MusicInfo, err
	}

	return ret.MusicInfo, nil
}

func (m MusicSearch) GetInfo() (MusicInfo, error) {
	return getInfo(m.MusicID)
}

func (m MusicSearchNew) GetInfo() (MusicInfo, error) {
	return getInfo(m.MusicID)
}

func (v VideoSearch) GetInfo() (MusicInfo, error) {
	return MusicInfo{}, nil
}

func (a ArtistSearch) GetInfo() (MusicInfo, error) {
	return MusicInfo{}, nil
}

func (a AlbumSearch) GetInfo() (MusicInfo, error) {
	return MusicInfo{}, nil
}

func (m MusicSearch) GetLink() string {
	return ""
}

func (m MusicSearchNew) GetLink() string {
	return m.MusicLink
}

func (m VideoSearch) GetLink() string {
	return m.VideoLink
}

func (m ArtistSearch) GetLink() string {
	return Home + m.ArtistLink
}

func (m AlbumSearch) GetLink() string {
	return Home + m.AlbumLink
}

func (m MusicSearch) GetID() interface{} {
	return m.MusicID
}

func (m MusicSearchNew) GetID() interface{} {
	return m.MusicID
}

func (m VideoSearch) GetID() interface{} {
	return m.VideoID
}

func (m ArtistSearch) GetID() interface{} {
	return m.ArtistID
}

func (m AlbumSearch) GetID() interface{} {
	return m.AlbumID
}
