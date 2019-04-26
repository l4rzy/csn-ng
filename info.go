/*
 * Copyright (C) 2019 l4rzy
 * MIT License
 */

package csn

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"errors"
)

func getInfo(id interface{}) (CSNMusicInfo, error) {
	var ret CSNMusicInfoResp
	surl := fmt.Sprintf(MUSIC_INFO_FMT, id)

	// get data
	resp, err := csn_client.Get(surl)
	if err != nil {
		return ret.MusicInfo, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("Responsed status code is not ok")
		return ret.MusicInfo, err
	}

	// read to body
	body, err := ioutil.ReadAll(resp.Body)
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

func (m CSNMusicSearch)GetInfo() (CSNMusicInfo, error) {
	return getInfo(m.MusicID)
}

func (m CSNMusicSearchNew)GetInfo() (CSNMusicInfo, error) {
	return getInfo(m.MusicID)
}

func (v CSNVideoSearch) GetInfo() (CSNMusicInfo, error) {
	return CSNMusicInfo{}, nil
}

func (a CSNArtistSearch) GetInfo() (CSNMusicInfo, error) {
	return CSNMusicInfo{}, nil
}

func (a CSNAlbumSearch) GetInfo() (CSNMusicInfo, error) {
	return CSNMusicInfo{}, nil
}
