package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	csn_client *http.Client
	err        error
)

const (
	search_fmt = "https://chiasenhac.vn/search/real?q=%s&type=json&rows=%d&view_all=true"
)

func init() {
	csn_client = &http.Client{Timeout: 10 * time.Second}
}

func buildSearchUrl(keyword string, limit int) (string, error) {
	if limit <= 0 || keyword == "" {
		err := errors.New("Wrong search input")
		return "", err
	}
	return fmt.Sprintf(search_fmt, keyword, limit), nil
}

func search(opt int, keyword string, limit int) ([]CSNObject, error) {
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

	// read to body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse json
	var result RespJson
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	var data = result[0]

	// return result as user asked
	var ret []CSNObject
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
