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
)

func (mi CSNMusicInfo) PrintLinks(opt int) {
	if opt&MUSIC_QUAL_32 != 0 {
		fmt.Printf("[ 32] %v\n", mi.File32URL)
	}
	if opt&MUSIC_QUAL_128 != 0 &&
		mi.FileURL != "" {
		fmt.Printf("[128] %v\n", mi.FileURL)
	}
	if opt&MUSIC_QUAL_320 != 0 &&
		mi.File320URL != "" {
		fmt.Printf("[320] %v\n", mi.File320URL)
	}
	if opt&MUSIC_QUAL_500 != 0 &&
		mi.FileM4AURL != "" {
		fmt.Printf("[500] %v\n", mi.FileM4AURL)
	}
	if opt&MUSIC_QUAL_1000 != 0 &&
		mi.FileLosslessURL != "" {
		fmt.Printf("[LLs] %v\n", mi.FileLosslessURL)
	}
}

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

func GetInfoUrl(url string) (CSNMusicInfo, error) {
	return CSNMusicInfo{}, nil
}

func (m CSNMusicSearch) GetInfo() (CSNMusicInfo, error) {
	return getInfo(m.MusicID)
}

func (m CSNMusicSearchNew) GetInfo() (CSNMusicInfo, error) {
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
